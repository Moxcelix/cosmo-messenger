package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	message_domain "main/internal/domain/message"
	"main/pkg"
)

type MessageRepository struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewMessageRepository(db pkg.PostgresDB, logger pkg.Logger) message_domain.MessageRepository {
	if err := createMessageTable(db); err != nil {
		logger.Error("failed to create messages table", err)
	}

	return &MessageRepository{
		db:     db,
		logger: logger,
	}
}

func createMessageTable(db pkg.PostgresDB) error {
	query := `
		CREATE TABLE IF NOT EXISTS messages (
			id VARCHAR PRIMARY KEY,
			chat_id VARCHAR NOT NULL,
			sender_id VARCHAR NOT NULL,
			reply_to VARCHAR DEFAULT '',
			content TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);

		CREATE TABLE IF NOT EXISTS message_attachments (
			id VARCHAR PRIMARY KEY,
			message_id VARCHAR NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
			type VARCHAR NOT NULL,
			url VARCHAR NOT NULL,
			filename VARCHAR,
			size BIGINT,
			mime_type VARCHAR,
			created_at TIMESTAMP NOT NULL
		);

		CREATE INDEX IF NOT EXISTS idx_messages_chat_id ON messages(chat_id);
		CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at DESC);
		CREATE INDEX IF NOT EXISTS idx_messages_sender_id ON messages(sender_id);
		CREATE INDEX IF NOT EXISTS idx_messages_reply_to ON messages(reply_to);
		CREATE INDEX IF NOT EXISTS idx_attachments_message_id ON message_attachments(message_id);
	`

	_, err := db.Exec(query)
	return err
}

func (r *MessageRepository) CreateMessage(message *message_domain.Message) error {
	if message.ID == "" {
		message.ID = generateMessageID()
	}

	now := time.Now()
	message.CreatedAt = now
	message.UpdatedAt = now

	return r.db.WithTransaction(context.Background(), func(tx *sql.Tx) error {
		query := `
			INSERT INTO messages (id, chat_id, sender_id, reply_to, content, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`

		_, err := tx.Exec(
			query,
			message.ID,
			message.ChatID,
			message.SenderID,
			message.ReplyTo,
			message.Content,
			message.CreatedAt,
			message.UpdatedAt,
		)
		if err != nil {
			return err
		}

		for _, attachment := range message.Attachments {
			attachmentQuery := `
				INSERT INTO message_attachments (id, message_id, type, url, filename, size, mime_type, created_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			`
			_, err = tx.Exec(
				attachmentQuery,
				attachment.ID,
				message.ID,
				attachment.Type,
				attachment.URL,
				attachment.Filename,
				attachment.Size,
				attachment.MimeType,
				now,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *MessageRepository) GetMessageById(id string) (*message_domain.Message, error) {
	messageQuery := `
		SELECT id, chat_id, sender_id, reply_to, content, created_at, updated_at 
		FROM messages 
		WHERE id = $1
	`

	var message message_domain.Message
	err := r.db.QueryRow(messageQuery, id).Scan(
		&message.ID,
		&message.ChatID,
		&message.SenderID,
		&message.ReplyTo,
		&message.Content,
		&message.CreatedAt,
		&message.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	attachments, err := r.getMessageAttachments(id)
	if err != nil {
		return nil, err
	}
	message.Attachments = attachments

	return &message, nil
}

func (r *MessageRepository) UpdateMessage(message *message_domain.Message) error {
	message.UpdatedAt = time.Now()

	return r.db.WithTransaction(context.Background(), func(tx *sql.Tx) error {
		query := `
			UPDATE messages 
			SET reply_to = $1, content = $2, updated_at = $3
			WHERE id = $4
		`

		result, err := tx.Exec(
			query,
			message.ReplyTo,
			message.Content,
			message.UpdatedAt,
			message.ID,
		)
		if err != nil {
			return err
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return message_domain.ErrMessageNotFound
		}

		_, err = tx.Exec("DELETE FROM message_attachments WHERE message_id = $1", message.ID)
		if err != nil {
			return err
		}

		for _, attachment := range message.Attachments {
			attachmentQuery := `
				INSERT INTO message_attachments (id, message_id, type, url, filename, size, mime_type, created_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			`
			_, err = tx.Exec(
				attachmentQuery,
				attachment.ID,
				message.ID,
				attachment.Type,
				attachment.URL,
				attachment.Filename,
				attachment.Size,
				attachment.MimeType,
				time.Now(),
			)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *MessageRepository) getMessageAttachments(messageID string) ([]message_domain.Attachment, error) {
	query := `
		SELECT id, type, url, filename, size, mime_type, created_at
		FROM message_attachments 
		WHERE message_id = $1
		ORDER BY created_at
	`

	rows, err := r.db.Query(query, messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []message_domain.Attachment
	for rows.Next() {
		var attachment message_domain.Attachment
		err := rows.Scan(
			&attachment.ID,
			&attachment.Type,
			&attachment.URL,
			&attachment.Filename,
			&attachment.Size,
			&attachment.MimeType,
			&attachment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, attachment)
	}

	return attachments, nil
}

func (r *MessageRepository) DeleteMessage(id string) error {
	query := `DELETE FROM messages WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return message_domain.ErrMessageNotFound
	}

	return nil
}

func (r *MessageRepository) GetLastChatMessage(chatId string) (*message_domain.Message, error) {
	query := `
        SELECT id, chat_id, sender_id, content, reply_to, created_at, updated_at
        FROM messages 
        WHERE chat_id = $1 
        ORDER BY created_at DESC 
        LIMIT 1
    `

	var message message_domain.Message
	err := r.db.QueryRow(query, chatId).Scan(
		&message.ID,
		&message.ChatID,
		&message.SenderID,
		&message.Content,
		&message.ReplyTo,
		&message.CreatedAt,
		&message.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	attachments, err := r.getMessageAttachments(message.ID)
	if err != nil {
		return nil, err
	}
	message.Attachments = attachments

	return &message, nil
}

func generateMessageID() string {
	return fmt.Sprintf("msg_%d", time.Now().UnixNano())
}
