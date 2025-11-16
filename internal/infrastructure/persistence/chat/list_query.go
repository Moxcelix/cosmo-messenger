package persistence

import (
	"context"
	"database/sql"
	"main/internal/application/chat/queries"
	"main/internal/application/chat/readmodels"
	"main/pkg"
)

type ChatListQuery struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewChatListQuery(db pkg.PostgresDB, logger pkg.Logger) queries.ChatListQuery {
	return &ChatListQuery{
		db:     db,
		logger: logger,
	}
}

func (q *ChatListQuery) Query(userID string, offset, limit int) (*readmodels.ChatList, error) {
	ctx := context.Background()

	query := `
        WITH UserChats AS (
            SELECT DISTINCT c.id as chat_id
            FROM chats c
            JOIN members m ON c.id = m.chat_id
            WHERE m.user_id = $1
        ),
        LastMessages AS (
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
            WHERE m.chat_id IN (SELECT chat_id FROM UserChats)
        )
        SELECT 
            c.id, c.type, c.name, c.description, c.updated_at,
            lm.message_id, lm.content, lm.reply_to, lm.created_at as message_created_at,
            lm.sender_id, lm.sender_name, lm.sender_username,
            (SELECT COUNT(*) FROM members m2 WHERE m2.chat_id = c.id) as member_count,
            (SELECT COUNT(*) FROM UserChats) as total_chats
        FROM chats c
        LEFT JOIN LastMessages lm ON c.id = lm.chat_id AND lm.rn = 1
        WHERE c.id IN (SELECT chat_id FROM UserChats)
        ORDER BY COALESCE(lm.created_at, c.updated_at) DESC
        LIMIT $2 OFFSET $3
    `

	rows, err := q.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*readmodels.ChatWithLastMessage
	var totalChats int

	for rows.Next() {
		var chat readmodels.ChatWithLastMessage
		var messageID, content, replyTo sql.NullString
		var messageCreatedAt sql.NullTime
		var senderID, senderName, senderUsername sql.NullString

		err := rows.Scan(
			&chat.ID, &chat.Type, &chat.Name, &chat.Description, &chat.UpdatedAt,
			&messageID, &content, &replyTo, &messageCreatedAt,
			&senderID, &senderName, &senderUsername,
			&chat.MemberCount, &totalChats,
		)
		if err != nil {
			return nil, err
		}

		if messageID.Valid {
			chat.LastMessage = &readmodels.LastMessage{
				ID:      messageID.String,
				Content: content.String,
				ReplyId: replyTo.String,
				SentAt:  messageCreatedAt.Time,
				Sender: &readmodels.LastMessageSender{
					ID:       senderID.String,
					Name:     senderName.String,
					Username: senderUsername.String,
				},
			}
		}

		chats = append(chats, &chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	hasNext := offset+len(chats) < totalChats
	hasPrev := offset > 0

	return &readmodels.ChatList{
		Chats:   chats,
		HasNext: hasNext,
		HasPrev: hasPrev,
	}, nil
}
