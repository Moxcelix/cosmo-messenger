package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	msg_errors "main/internal_new/domain/chat/errors"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/repositories"
	"main/pkg"
	"time"
)

type PostgresMessageRepository struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewPostgresMessageRepository(db pkg.PostgresDB, logger pkg.Logger) repositories.MessageRepository {
	if err := createMessageTable(db); err != nil {
		logger.Error("failed to create messages table", err)
	}

	return &PostgresMessageRepository{
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

func (r *PostgresMessageRepository) CreateMessage(message *models.Message) error {
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
			message.ReplyToId,
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

func (r *PostgresMessageRepository) GetMessageById(id string) (*models.Message, error) {
	messageQuery := `
		SELECT id, chat_id, sender_id, reply_to, content, created_at, updated_at 
		FROM messages 
		WHERE id = $1
	`

	var message models.Message
	err := r.db.QueryRow(messageQuery, id).Scan(
		&message.ID,
		&message.ChatID,
		&message.SenderID,
		&message.ReplyToId,
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

func (r *PostgresMessageRepository) UpdateMessage(message *models.Message) error {
	message.UpdatedAt = time.Now()

	return r.db.WithTransaction(context.Background(), func(tx *sql.Tx) error {
		query := `
			UPDATE messages 
			SET reply_to = $1, content = $2, updated_at = $3
			WHERE id = $4
		`

		result, err := tx.Exec(
			query,
			message.ReplyToId,
			message.Content,
			message.UpdatedAt,
			message.ID,
		)
		if err != nil {
			return err
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return msg_errors.ErrMessageNotFound
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

func (r *PostgresMessageRepository) getMessageAttachments(messageID string) ([]models.Attachment, error) {
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

	var attachments []models.Attachment
	for rows.Next() {
		var attachment models.Attachment
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

func (r *PostgresMessageRepository) DeleteMessage(id string) error {
	query := `DELETE FROM messages WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return msg_errors.ErrMessageNotFound
	}

	return nil
}

func (r *PostgresMessageRepository) GetLastChatMessage(chatId string) (*models.Message, error) {
	query := `
        SELECT id, chat_id, sender_id, content, reply_to, created_at, updated_at
        FROM messages 
        WHERE chat_id = $1 
        ORDER BY created_at DESC 
        LIMIT 1
    `

	var message models.Message
	err := r.db.QueryRow(query, chatId).Scan(
		&message.ID,
		&message.ChatID,
		&message.SenderID,
		&message.Content,
		&message.ReplyToId,
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
