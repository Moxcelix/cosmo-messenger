package chat_infrastructure

import (
	"main/internal/config"
	chat_domain "main/internal/domain/chat"
	"sync"
	"time"
)

type InMemoryTypingService struct {
	sessions        map[string]map[string]*chat_domain.TypingSession
	mutex           sync.RWMutex
	cleanupInterval time.Duration
	ttl             time.Duration
}

func NewInMemoryTypingService(env config.Env) *InMemoryTypingService {
	return &InMemoryTypingService{
		sessions:        make(map[string]map[string]*chat_domain.TypingSession),
		cleanupInterval: env.TypingCleanupInterval,
		ttl:             env.TypingTTL,
	}
}

func (s *InMemoryTypingService) StartTyping(userID, chatID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.sessions[chatID] == nil {
		s.sessions[chatID] = make(map[string]*chat_domain.TypingSession)
	}

	if session, exists := s.sessions[chatID][userID]; exists {
		session.Update()
	} else {
		s.sessions[chatID][userID] = chat_domain.NewTypingSession(userID, chatID)
	}

	return nil
}

func (s *InMemoryTypingService) StopTyping(userID, chatID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if chatSessions, exists := s.sessions[chatID]; exists {
		delete(chatSessions, userID)

		if len(chatSessions) == 0 {
			delete(s.sessions, chatID)
		}
	}

	return nil
}

func (s *InMemoryTypingService) GetActiveSessions(chatID string) ([]*chat_domain.TypingSession, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var activeSessions []*chat_domain.TypingSession

	if chatSessions, exists := s.sessions[chatID]; exists {
		for _, session := range chatSessions {
			activeSessions = append(activeSessions, session)
		}
	}

	return activeSessions, nil
}

func (s *InMemoryTypingService) GetActiveUsers(chatID string, ttl time.Duration) ([]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var activeUsers []string

	if chatSessions, exists := s.sessions[chatID]; exists {
		for userID, session := range chatSessions {
			if !session.IsExpired(ttl) {
				activeUsers = append(activeUsers, userID)
			}
		}
	}

	return activeUsers, nil
}

func (s *InMemoryTypingService) CleanupExpiredSessions(ttl time.Duration) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for chatID, chatSessions := range s.sessions {
		for userID, session := range chatSessions {
			if session.IsExpired(ttl) {
				delete(chatSessions, userID)
			}
		}

		if len(chatSessions) == 0 {
			delete(s.sessions, chatID)
		}
	}

	return nil
}

func (s *InMemoryTypingService) GetUserSession(userID, chatID string) (*chat_domain.TypingSession, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if chatSessions, exists := s.sessions[chatID]; exists {
		if session, exists := chatSessions[userID]; exists {
			return session, true
		}
	}

	return nil, false
}

func (s *InMemoryTypingService) StartCleanup(interval, ttl time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			s.CleanupExpiredSessions(ttl)
		}
	}()
}

func (s *InMemoryTypingService) Run() {
	s.StartCleanup(s.cleanupInterval, s.ttl)
}
