package message_infrastructure

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

		CREATE INDEX IF NOT EXISTS idx_messages_chat_id ON messages(chat_id);
		CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at DESC);
		CREATE INDEX IF NOT EXISTS idx_messages_sender_id ON messages(sender_id);
		CREATE INDEX IF NOT EXISTS idx_messages_reply_to ON messages(reply_to);
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

	query := `
		INSERT INTO messages (id, chat_id, sender_id, reply_to, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(
		query,
		message.ID,
		message.ChatID,
		message.SenderID,
		message.ReplyTo,
		message.Content,
		message.CreatedAt,
		message.UpdatedAt,
	)

	return err
}

func (r *MessageRepository) GetMessageById(id string) (*message_domain.Message, error) {
	query := `
		SELECT id, chat_id, sender_id, reply_to, content, created_at, updated_at 
		FROM messages 
		WHERE id = $1
	`

	var message message_domain.Message
	err := r.db.QueryRow(query, id).Scan(
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

	return &message, nil
}

func (r *MessageRepository) UpdateMessage(message *message_domain.Message) error {
	message.UpdatedAt = time.Now()

	query := `
		UPDATE messages 
		SET reply_to = $1, content = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(
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

	return nil
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

func (r *MessageRepository) GetMessagesByChatId(
	chatId string, offset, limit int) (*message_domain.MessageList, error) {

	ctx := context.Background()

	var total int
	countQuery := `SELECT COUNT(*) FROM messages WHERE chat_id = $1`
	err := r.db.QueryRowContext(ctx, countQuery, chatId).Scan(&total)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, chat_id, sender_id, reply_to, content, created_at, updated_at 
		FROM messages 
		WHERE chat_id = $1
		ORDER BY created_at DESC 
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, chatId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*message_domain.Message
	for rows.Next() {
		var message message_domain.Message
		err := rows.Scan(
			&message.ID,
			&message.ChatID,
			&message.SenderID,
			&message.ReplyTo,
			&message.Content,
			&message.CreatedAt,
			&message.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &message_domain.MessageList{
		Messages: messages,
		Total:    total,
		Offset:   offset,
		Limit:    limit,
	}, nil
}

func generateMessageID() string {
	return fmt.Sprintf("msg_%d", time.Now().UnixNano())
}
