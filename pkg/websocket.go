package pkg

import (
	"fmt"
	"main/internal/config"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connections map[*websocket.Conn]ConnectionPort
	mutex       sync.RWMutex
}

type ConnectionPort struct {
	ConnectedAt time.Time
	UserAgent   string
	IPAddress   string
	stopPing    chan struct{}
}

type WebSocketEvent struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type MessageHandler func(clientID string, message WebSocketEvent) error

type WebSocketHub struct {
	upgrader      websocket.Upgrader
	clients       map[string]*Client
	clientsMutex  sync.RWMutex
	handlers      map[string]MessageHandler
	handlersMutex sync.RWMutex
	logger        Logger
}

func NewWebSocketHub(env config.Env, logger Logger) *WebSocketHub {
	return &WebSocketHub{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				if env.AppEnv != "production" {
					return true
				}
				return slices.Contains(env.AllowedOrigins, origin)
			},
		},
		clients:  make(map[string]*Client),
		handlers: make(map[string]MessageHandler),
		logger:   logger,
	}
}

func (h *WebSocketHub) HandleConnection(w http.ResponseWriter, r *http.Request, clientID string) error {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	connectionInfo := ConnectionPort{
		ConnectedAt: time.Now(),
		UserAgent:   r.UserAgent(),
		IPAddress:   r.RemoteAddr,
		stopPing:    make(chan struct{}),
	}

	h.clientsMutex.Lock()
	if _, exists := h.clients[clientID]; !exists {
		h.clients[clientID] = &Client{
			Connections: make(map[*websocket.Conn]ConnectionPort),
		}
	}

	client := h.clients[clientID]
	client.mutex.Lock()
	client.Connections[conn] = connectionInfo
	client.mutex.Unlock()
	h.clientsMutex.Unlock()

	h.logger.Info("Client connected", "clientID", clientID, "ip", connectionInfo.IPAddress)

	go h.handleConnection(conn, clientID, connectionInfo.stopPing)

	return nil
}

func (h *WebSocketHub) handleConnection(conn *websocket.Conn, clientID string, stopPing chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			h.logger.Error("WebSocket panic in handleConnection", "error", err, "clientID", clientID)
		}
	}()

	defer func() {
		h.RemoveConnection(clientID, conn)
		conn.Close()
	}()

	messageChan := make(chan WebSocketEvent)
	errorChan := make(chan error)
	done := make(chan struct{})

	go h.readMessages(conn, messageChan, errorChan, done)

	pingTicker := time.NewTicker(30 * time.Second)
	defer pingTicker.Stop()

	for {
		select {
		case event := <-messageChan:
			h.handleEvent(event, clientID)

		case err := <-errorChan:
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				h.logger.Error("WebSocket error", "error", err, "clientID", clientID)
			}
			return

		case <-pingTicker.C:
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				h.logger.Error("Ping failed", "clientID", clientID, "error", err)
				return
			}

		case <-stopPing:
			return

		case <-done:
			return
		}
	}
}

func (h *WebSocketHub) readMessages(conn *websocket.Conn, messageChan chan<- WebSocketEvent, errorChan chan<- error, done chan<- struct{}) {
	defer close(done)

	for {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		var event WebSocketEvent
		err := conn.ReadJSON(&event)
		if err != nil {
			errorChan <- err
			return
		}

		messageChan <- event
	}
}

func (h *WebSocketHub) handleEvent(event WebSocketEvent, clientID string) {
	h.handlersMutex.RLock()
	handler, exists := h.handlers[event.Type]
	h.handlersMutex.RUnlock()

	if !exists {
		h.logger.Warn("No handler for event type", "type", event.Type, "clientID", clientID)
		return
	}

	if err := handler(clientID, event); err != nil {
		h.logger.Error("Handler error", "error", err, "clientID", clientID, "type", event.Type)
	}
}

func (h *WebSocketHub) On(eventType string, handler MessageHandler) {
	h.handlersMutex.Lock()
	defer h.handlersMutex.Unlock()
	h.handlers[eventType] = handler
}

func (h *WebSocketHub) SendToClient(clientID string, event WebSocketEvent) error {
	h.clientsMutex.RLock()
	client, exists := h.clients[clientID]
	h.clientsMutex.RUnlock()

	if !exists {
		return fmt.Errorf("client not found: %s", clientID)
	}

	client.mutex.RLock()
	defer client.mutex.RUnlock()

	var errors []error
	for conn := range client.Connections {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(event)
		if err != nil {
			errors = append(errors, err)
			h.logger.Error("Send error", "clientID", clientID, "error", err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("errors sending to client %s: %v", clientID, errors)
	}

	return nil
}

func (h *WebSocketHub) Broadcast(event WebSocketEvent) {
	h.clientsMutex.RLock()
	defer h.clientsMutex.RUnlock()

	for clientID, client := range h.clients {
		client.mutex.RLock()
		for conn := range client.Connections {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteJSON(event); err != nil {
				h.logger.Error("Broadcast error", "error", err, "clientID", clientID)
			}
		}
		client.mutex.RUnlock()
	}
}

func (h *WebSocketHub) RemoveConnection(clientID string, conn *websocket.Conn) {
	h.clientsMutex.Lock()
	defer h.clientsMutex.Unlock()

	client, exists := h.clients[clientID]
	if !exists {
		return
	}

	client.mutex.Lock()
	if port, exists := client.Connections[conn]; exists {
		if port.stopPing != nil {
			close(port.stopPing)
		}
		delete(client.Connections, conn)
	}
	client.mutex.Unlock()

	if len(client.Connections) == 0 {
		delete(h.clients, clientID)
		h.logger.Info("Client removed", "clientID", clientID)
	}
}

func (h *WebSocketHub) RemoveClient(clientID string) {
	h.clientsMutex.Lock()
	defer h.clientsMutex.Unlock()

	client, exists := h.clients[clientID]
	if !exists {
		return
	}

	client.mutex.Lock()
	for conn, port := range client.Connections {
		if port.stopPing != nil {
			close(port.stopPing)
		}
		conn.Close()
	}
	client.mutex.Unlock()

	delete(h.clients, clientID)
	h.logger.Info("Client fully removed", "clientID", clientID)
}

func (h *WebSocketHub) GetClient(clientID string) (*Client, bool) {
	h.clientsMutex.RLock()
	defer h.clientsMutex.RUnlock()

	client, exists := h.clients[clientID]
	return client, exists
}

func (h *WebSocketHub) GetClientCount() int {
	h.clientsMutex.RLock()
	defer h.clientsMutex.RUnlock()

	return len(h.clients)
}

func (h *WebSocketHub) GetConnectionCount() int {
	h.clientsMutex.RLock()
	defer h.clientsMutex.RUnlock()

	count := 0
	for _, client := range h.clients {
		client.mutex.RLock()
		count += len(client.Connections)
		client.mutex.RUnlock()
	}

	return count
}
