package persistence

import (
	"context"
	"database/sql"
	"main/internal/application/message/queries"
	"main/internal/application/message/readmodels"
	user_readmodels "main/internal/application/user/readmodels"
	"main/pkg"
	"time"
)

type MessageQuery struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewMessageQuery(db pkg.PostgresDB, logger pkg.Logger) queries.MessageQuery {
	return &MessageQuery{
		db:     db,
		logger: logger,
	}
}

func (q *MessageQuery) Query(msgId string) (*readmodels.Message, error) {
	ctx := context.Background()

	query := `
		SELECT 
			m.id,
			m.content,
			m.chat_id,
			m.reply_to,
			m.created_at,
			m.updated_at,
			-- sender data
			u.id as sender_id,
			u.name as sender_name,
			u.username as sender_username,
			-- reply data
			rm.content as reply_content,
			-- reply sender data
			ru.id as reply_sender_id,
			ru.name as reply_sender_name,
			ru.username as reply_sender_username
		FROM messages m
		INNER JOIN users u ON m.sender_id = u.id
		LEFT JOIN messages rm ON m.reply_to = rm.id
		LEFT JOIN users ru ON rm.sender_id = ru.id
		WHERE m.id = $1
	`

	var (
		messageID, content, chatID, replyToID                             sql.NullString
		createdAt, updatedAt                                              time.Time
		senderID, senderName, senderUsername                              string
		replyContent, replySenderID, replySenderName, replySenderUsername sql.NullString
	)

	err := q.db.QueryRowContext(ctx, query, msgId).Scan(
		&messageID, &content, &chatID, &replyToID, &createdAt, &updatedAt,
		&senderID, &senderName, &senderUsername,
		&replyContent, &replySenderID, &replySenderName, &replySenderUsername,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	msg := &readmodels.Message{
		ID:        messageID.String,
		Content:   content.String,
		ChatID:    chatID.String,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Sender: &user_readmodels.Sender{
			ID:       senderID,
			Name:     senderName,
			Username: senderUsername,
		},
		Attachments: []*readmodels.Attachment{},
	}

	if replyToID.Valid && replyContent.Valid {
		msg.ReplyTo = &readmodels.Reply{
			ID:      replyToID.String,
			Content: replyContent.String,
		}

		if replySenderID.Valid && replySenderName.Valid {
			msg.ReplyTo.Sender = &user_readmodels.Sender{
				ID:       replySenderID.String,
				Name:     replySenderName.String,
				Username: replySenderUsername.String,
			}
		}
	}

	err = q.loadAttachments(ctx, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (q *MessageQuery) loadAttachments(ctx context.Context, message *readmodels.Message) error {
	query := `
		SELECT 
			id, type, url, filename, size, mime_type
		FROM message_attachments 
		WHERE message_id = $1
		ORDER BY created_at
	`

	rows, err := q.db.QueryContext(ctx, query, message.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var attachments []*readmodels.Attachment

	for rows.Next() {
		var attachment readmodels.Attachment

		err := rows.Scan(
			&attachment.ID,
			&attachment.Type,
			&attachment.URL,
			&attachment.Filename,
			&attachment.Size,
			&attachment.MimeType,
		)
		if err != nil {
			return err
		}

		attachments = append(attachments, &attachment)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	message.Attachments = attachments
	return nil
}
