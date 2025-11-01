package chat_domain

import "time"

type TypingSession struct {
	UserID    string
	ChatID    string
	StartedAt time.Time
	UpdatedAt time.Time
}

func NewTypingSession(userID, chatID string) *TypingSession {
	now := time.Now()
	return &TypingSession{
		UserID:    userID,
		ChatID:    chatID,
		StartedAt: now,
		UpdatedAt: now,
	}
}

func (ts *TypingSession) IsExpired(ttl time.Duration) bool {
	return time.Since(ts.UpdatedAt) > ttl
}

func (ts *TypingSession) Update() {
	ts.UpdatedAt = time.Now()
}

func (ts *TypingSession) GetDuration() time.Duration {
	return time.Since(ts.StartedAt)
}

type TypingService interface {
	StartTyping(userID, chatID string) error
	StopTyping(userID, chatID string) error
	GetActiveSessions(chatID string) ([]*TypingSession, error)
	CleanupExpiredSessions(ttl time.Duration) error
}
