package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/application/message/queries"
	"main/internal/application/message/readmodels"
	"main/pkg"
	"strings"
	"time"

	user_readmodels "main/internal/application/user/readmodels"
)

type MessageHistoryQuery struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewMessageHistoryQuery(db pkg.PostgresDB, logger pkg.Logger) queries.MessageHistoryQuery {
	return &MessageHistoryQuery{
		db:     db,
		logger: logger,
	}
}

func (q *MessageHistoryQuery) Query(chatId string, cursorMessageId string, count int, direction string) (*readmodels.MessageHistory, error) {
	ctx := context.Background()

	chatHeader, err := q.getChatHeader(ctx, chatId)
	if err != nil {
		return nil, err
	}

	if chatHeader == nil {
		return &readmodels.MessageHistory{
			ChatHeader: nil,
			Messages:   []*readmodels.Message{},
			HasNext:    false,
			HasPrev:    false,
		}, nil
	}

	messages, hasNext, hasPrev, err := q.getMessagesWithPagination(ctx, chatId, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	return &readmodels.MessageHistory{
		ChatHeader: chatHeader,
		Messages:   messages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}, nil
}

func (q *MessageHistoryQuery) getChatHeader(ctx context.Context, chatId string) (*readmodels.ChatHeader, error) {
	query := `
		SELECT id, type, name 
		FROM chats 
		WHERE id = $1
	`

	var header readmodels.ChatHeader
	err := q.db.QueryRowContext(ctx, query, chatId).Scan(
		&header.ID,
		&header.Type,
		&header.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &header, nil
}

func (q *MessageHistoryQuery) getMessagesWithPagination(
	ctx context.Context,
	chatId string,
	cursorMessageId string,
	count int,
	direction string,
) ([]*readmodels.Message, bool, bool, error) {

	messages, err := q.getMessagesOnly(ctx, chatId, cursorMessageId, count, direction)
	if err != nil {
		return nil, false, false, err
	}

	if len(messages) == 0 {
		return messages, false, false, nil
	}

	err = q.loadAttachments(ctx, messages)
	if err != nil {
		return nil, false, false, err
	}

	hasNext, hasPrev, err := q.checkPaginationBounds(ctx, chatId, messages)
	if err != nil {
		return nil, false, false, err
	}

	return messages, hasNext, hasPrev, nil
}

func (q *MessageHistoryQuery) getMessagesOnly(
	ctx context.Context,
	chatId string,
	cursorMessageId string,
	count int,
	direction string,
) ([]*readmodels.Message, error) {

	baseQuery := `
		SELECT 
			m.id,
			m.content,
			m.chat_id,
			m.reply_to,
			m.created_at,
			m.updated_at,
			-- Данные отправителя
			u.id as sender_id,
			u.name as sender_name,
			u.username as sender_username,
			-- Данные реплая (если есть)
			rm.content as reply_content,
			ru.id as reply_sender_id,
			ru.name as reply_sender_name,
			ru.username as reply_sender_username
		FROM messages m
		INNER JOIN users u ON m.sender_id = u.id
		LEFT JOIN messages rm ON m.reply_to = rm.id
		LEFT JOIN users ru ON rm.sender_id = ru.id
		WHERE m.chat_id = $1
	`

	var query string
	var args []interface{}

	switch direction {
	case "older":
		if cursorMessageId != "" {
			query = baseQuery + ` AND m.id < $2 ORDER BY m.id DESC LIMIT $3`
			args = []interface{}{chatId, cursorMessageId, count}
		} else {
			query = baseQuery + ` ORDER BY m.id DESC LIMIT $2`
			args = []interface{}{chatId, count}
		}

	case "newer":
		if cursorMessageId != "" {
			query = baseQuery + ` AND m.id > $2 ORDER BY m.id ASC LIMIT $3`
			args = []interface{}{chatId, cursorMessageId, count}
		} else {
			query = baseQuery + ` ORDER BY m.id ASC LIMIT $2`
			args = []interface{}{chatId, count}
		}

	default:
		query = baseQuery + ` ORDER BY m.id DESC LIMIT $2`
		args = []interface{}{chatId, count}
	}

	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*readmodels.Message

	for rows.Next() {
		var (
			messageID, content, chatID, replyToID                             sql.NullString
			createdAt, updatedAt                                              time.Time
			senderID, senderName, senderUsername                              string
			replyContent, replySenderID, replySenderName, replySenderUsername sql.NullString
		)

		err := rows.Scan(
			&messageID, &content, &chatID, &replyToID, &createdAt, &updatedAt,
			&senderID, &senderName, &senderUsername,
			&replyContent, &replySenderID, &replySenderName, &replySenderUsername,
		)
		if err != nil {
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

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if direction == "newer" {
		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages[i], messages[j] = messages[j], messages[i]
		}
	}

	return messages, nil
}

func (q *MessageHistoryQuery) loadAttachments(ctx context.Context, messages []*readmodels.Message) error {
	if len(messages) == 0 {
		return nil
	}

	placeholders := make([]string, len(messages))
	args := make([]interface{}, len(messages))

	for i, msg := range messages {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = msg.ID
	}

	query := fmt.Sprintf(`
        SELECT 
            message_id,
            id, type, url, filename, size, mime_type
        FROM message_attachments 
        WHERE message_id IN (%s)
        ORDER BY created_at
    `, strings.Join(placeholders, ","))

	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	attachmentsByMessage := make(map[string][]*readmodels.Attachment)

	for rows.Next() {
		var (
			messageID  string
			attachment readmodels.Attachment
		)

		err := rows.Scan(
			&messageID,
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

		attachmentsByMessage[messageID] = append(attachmentsByMessage[messageID], &attachment)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for _, msg := range messages {
		if atts, exists := attachmentsByMessage[msg.ID]; exists {
			msg.Attachments = atts
		}
	}

	return nil
}

func (q *MessageHistoryQuery) checkPaginationBounds(
	ctx context.Context,
	chatId string,
	messages []*readmodels.Message,
) (bool, bool, error) {

	if len(messages) == 0 {
		return false, false, nil
	}

	var hasNext, hasPrev bool
	var err error

	newestMessage := messages[0]
	oldestMessage := messages[len(messages)-1]

	hasPrevQuery := `SELECT EXISTS(SELECT 1 FROM messages WHERE chat_id = $1 AND id < $2)`
	err = q.db.QueryRowContext(ctx, hasPrevQuery, chatId, oldestMessage.ID).Scan(&hasPrev)
	if err != nil {
		return false, false, err
	}

	hasNextQuery := `SELECT EXISTS(SELECT 1 FROM messages WHERE chat_id = $1 AND id > $2)`
	err = q.db.QueryRowContext(ctx, hasNextQuery, chatId, newestMessage.ID).Scan(&hasNext)
	if err != nil {
		return false, false, err
	}

	return hasNext, hasPrev, nil
}
