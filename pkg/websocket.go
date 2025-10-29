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
	Connections map[*websocket.Conn]ConnectionInfo

	mutex sync.RWMutex
}

type ConnectionInfo struct {
	ConnectedAt time.Time
	UserAgent   string
	IPAddress   string
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

func (m *WebSocketHub) HandleConnection(w http.ResponseWriter, r *http.Request, clientID string) error {
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	connectionInfo := ConnectionInfo{
		ConnectedAt: time.Now(),
		UserAgent:   r.UserAgent(),
		IPAddress:   r.RemoteAddr,
	}

	m.clientsMutex.Lock()
	defer m.clientsMutex.Unlock()

	if _, exists := m.clients[clientID]; !exists {
		m.clients[clientID] = &Client{
			Connections: make(map[*websocket.Conn]ConnectionInfo),
			mutex:       sync.RWMutex{},
		}
	}

	client := m.clients[clientID]
	client.mutex.Lock()
	client.Connections[conn] = connectionInfo
	client.mutex.Unlock()

	m.logger.Info("Client connected", "clientID", clientID, "ip", connectionInfo.IPAddress)

	go m.handleMessages(conn, clientID)

	return nil
}

func (m *WebSocketHub) handleMessages(conn *websocket.Conn, clientID string) {
	defer func() {
		m.RemoveConnection(clientID, conn)
		conn.Close()
	}()

	for {
		var event WebSocketEvent
		err := conn.ReadJSON(&event)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				m.logger.Error("WebSocket error", "error", err, "clientID", clientID)
			}
			break
		}

		m.handleEvent(event, clientID)
	}
}

func (m *WebSocketHub) handleEvent(event WebSocketEvent, clientID string) {
	m.handlersMutex.RLock()
	handler, exists := m.handlers[event.Type]
	m.handlersMutex.RUnlock()

	if !exists {
		m.logger.Warn("No handler for event type", "type", event.Type, "clientID", clientID)
		return
	}

	if err := handler(clientID, event); err != nil {
		m.logger.Error("Handler error", "error", err, "clientID", clientID, "type", event.Type)
	}
}

func (m *WebSocketHub) On(eventType string, handler MessageHandler) {
	m.handlersMutex.Lock()
	defer m.handlersMutex.Unlock()
	m.handlers[eventType] = handler
}

func (m *WebSocketHub) SendToClient(clientID string, event WebSocketEvent) error {
	m.clientsMutex.RLock()
	client, exists := m.clients[clientID]
	m.clientsMutex.RUnlock()

	if !exists {
		return fmt.Errorf("client not found: %s", clientID)
	}

	client.mutex.RLock()
	defer client.mutex.RUnlock()

	var errors []error
	for conn := range client.Connections {
		err := conn.WriteJSON(event)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("errors sending to client %s: %v", clientID, errors)
	}

	return nil
}

func (m *WebSocketHub) Broadcast(event WebSocketEvent) {
	m.clientsMutex.RLock()
	defer m.clientsMutex.RUnlock()

	for clientID, client := range m.clients {
		client.mutex.RLock()
		for conn := range client.Connections {
			if err := conn.WriteJSON(event); err != nil {
				m.logger.Error("Broadcast error", "error", err, "clientID", clientID)
			}
		}
		client.mutex.RUnlock()
	}
}

func (m *WebSocketHub) RemoveConnection(clientID string, conn *websocket.Conn) {
	m.clientsMutex.Lock()
	defer m.clientsMutex.Unlock()

	client, exists := m.clients[clientID]
	if !exists {
		return
	}

	client.mutex.Lock()
	delete(client.Connections, conn)
	client.mutex.Unlock()

	if len(client.Connections) == 0 {
		delete(m.clients, clientID)
		m.logger.Info("Client removed", "clientID", clientID)
	}
}

func (m *WebSocketHub) GetClient(clientID string) (*Client, bool) {
	m.clientsMutex.RLock()
	defer m.clientsMutex.RUnlock()

	client, exists := m.clients[clientID]
	return client, exists
}

func (m *WebSocketHub) GetClientCount() int {
	m.clientsMutex.RLock()
	defer m.clientsMutex.RUnlock()

	return len(m.clients)
}

func (m *WebSocketHub) GetConnectionCount() int {
	m.clientsMutex.RLock()
	defer m.clientsMutex.RUnlock()

	count := 0
	for _, client := range m.clients {
		client.mutex.RLock()
		count += len(client.Connections)
		client.mutex.RUnlock()
	}

	return count
}
