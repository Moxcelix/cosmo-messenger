package message_test

import (
	message_domain "main/internal/domain/message"
	"testing"
	"time"
)

type mockConfig struct {
	editDuration   time.Duration
	deleteDuration time.Duration
	minLength      int
	maxLength      int
}

func (m mockConfig) EditDuration() time.Duration   { return m.editDuration }
func (m mockConfig) DeleteDuration() time.Duration { return m.deleteDuration }
func (m mockConfig) MaxLength() int                { return m.maxLength }
func (m mockConfig) MinLength() int                { return m.minLength }

func TestPolicy(t *testing.T) {
	cfg := mockConfig{
		editDuration:   10 * time.Minute,
		deleteDuration: 15 * time.Minute,
		minLength:      3,
		maxLength:      10,
	}

	policy := message_domain.NewMessagePolicy(cfg)

	t.Run("ValidateMessageContent", func(t *testing.T) {
		tests := []struct {
			name    string
			message string
			wantErr bool
		}{
			{"too short", "hi", true},
			{"too long", "hello world", true},
			{"ok", "hello", false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := policy.ValidateMessageContent(tt.message)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateMessageContent(%q) error = %v, wantErr %v", tt.message, err, tt.wantErr)
				}
			})
		}
	})

	t.Run("ValidateDelete", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			name    string
			created time.Time
			wantErr bool
		}{
			{"within delete duration", now.Add(-10 * time.Minute), false},
			{"after delete duration", now.Add(-20 * time.Minute), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				msg := message_domain.Message{CreatedAt: tt.created}
				err := policy.ValidateDelete(msg, now)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateDelete(%v) error = %v, wantErr %v", msg.CreatedAt, err, tt.wantErr)
				}
			})
		}
	})

	t.Run("ValidateEdit", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			name    string
			created time.Time
			wantErr bool
		}{
			{"within edit duration", now.Add(-5 * time.Minute), false},
			{"after edit duration", now.Add(-15 * time.Minute), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				msg := message_domain.Message{CreatedAt: tt.created}
				err := policy.ValidateEdit(msg, now)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateEdit(%v) error = %v, wantErr %v", msg.CreatedAt, err, tt.wantErr)
				}
			})
		}
	})
}
