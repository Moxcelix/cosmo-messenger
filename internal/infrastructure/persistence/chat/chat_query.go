package persistence

import (
	"context"
	"database/sql"
	"main/internal/application/chat/queries"
	"main/internal/application/chat/readmodels"
	"main/pkg"

	user_readmodels "main/internal/application/user/readmodels"
)

type ChatQuery struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewChatQuery(db pkg.PostgresDB, logger pkg.Logger) queries.ChatQuery {
	return &ChatQuery{
		db:     db,
		logger: logger,
	}
}

func (q *ChatQuery) Query(chatID, userID string) (*readmodels.ChatWithLastMessage, error) {
	ctx := context.Background()

	query := `
        WITH LastMessages AS (
            SELECT 
                m.chat_id,
                m.id as message_id,
                m.content,
                m.reply_to,
                m.created_at,
                m.sender_id,
                u.name as sender_name,
                u.username as sender_username,
                ROW_NUMBER() OVER (PARTITION BY m.chat_id ORDER BY m.created_at DESC) as rn
            FROM messages m
            JOIN users u ON m.sender_id = u.id
            WHERE m.chat_id = $1
        )
        SELECT 
            c.id, c.type, c.name, c.description, c.updated_at,
            lm.message_id, lm.content, lm.reply_to, lm.created_at as message_created_at,
            lm.sender_id, lm.sender_name, lm.sender_username,
            (SELECT COUNT(*) FROM members m2 WHERE m2.chat_id = c.id) as member_count
        FROM chats c
        LEFT JOIN LastMessages lm ON c.id = lm.chat_id AND lm.rn = 1
        WHERE c.id = $1 
        AND EXISTS (
            SELECT 1 FROM members m 
            WHERE m.chat_id = c.id AND m.user_id = $2
        )
    `

	var chat readmodels.ChatWithLastMessage
	var messageID, content, replyTo sql.NullString
	var messageCreatedAt sql.NullTime
	var senderID, senderName, senderUsername sql.NullString

	err := q.db.QueryRowContext(ctx, query, chatID, userID).Scan(
		&chat.ID, &chat.Type, &chat.Name, &chat.Description, &chat.UpdatedAt,
		&messageID, &content, &replyTo, &messageCreatedAt,
		&senderID, &senderName, &senderUsername,
		&chat.MemberCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if messageID.Valid {
		chat.LastMessage = &readmodels.LastMessage{
			ID:      messageID.String,
			Content: content.String,
			ReplyId: replyTo.String,
			SentAt:  messageCreatedAt.Time,
			Sender: &user_readmodels.Sender{
				ID:       senderID.String,
				Name:     senderName.String,
				Username: senderUsername.String,
			},
		}
	}

	return &chat, nil
}
