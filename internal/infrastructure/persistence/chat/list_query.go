package persistence

import (
	"context"
	"database/sql"
	"main/internal/application/chat/queries"
	"main/internal/application/chat/readmodels"
	"main/pkg"

	user_readmodels "main/internal/application/user/readmodels"
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

func (q *ChatListQuery) Query(userID string, cursorChatID string, count int, direction string) (*readmodels.ChatList, error) {
	ctx := context.Background()

	baseQuery := `
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
            (SELECT COUNT(*) FROM members m2 WHERE m2.chat_id = c.id) as member_count
        FROM chats c
        LEFT JOIN LastMessages lm ON c.id = lm.chat_id AND lm.rn = 1
        WHERE c.id IN (SELECT chat_id FROM UserChats)
    `

	var query string
	var args []interface{}

	switch direction {
	case "older":
		if cursorChatID != "" {
			cursorTime, err := q.getChatUpdatedAt(ctx, cursorChatID)
			if err != nil {
				return nil, err
			}
			query = baseQuery + ` AND (c.updated_at < $2 OR (c.updated_at = $2 AND c.id < $3)) 
                               ORDER BY c.updated_at DESC, c.id DESC LIMIT $4`
			args = []interface{}{userID, cursorTime, cursorChatID, count}
		} else {
			query = baseQuery + ` ORDER BY c.updated_at DESC, c.id DESC LIMIT $2`
			args = []interface{}{userID, count}
		}

	case "newer":
		if cursorChatID != "" {
			cursorTime, err := q.getChatUpdatedAt(ctx, cursorChatID)
			if err != nil {
				return nil, err
			}
			query = baseQuery + ` AND (c.updated_at > $2 OR (c.updated_at = $2 AND c.id > $3)) 
                               ORDER BY c.updated_at ASC, c.id ASC LIMIT $4`
			args = []interface{}{userID, cursorTime, cursorChatID, count}
		} else {
			query = baseQuery + ` ORDER BY c.updated_at ASC, c.id ASC LIMIT $2`
			args = []interface{}{userID, count}
		}

	default:
		query = baseQuery + ` ORDER BY c.updated_at DESC, c.id DESC LIMIT $2`
		args = []interface{}{userID, count}
	}

	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*readmodels.ChatWithLastMessage

	for rows.Next() {
		var chat readmodels.ChatWithLastMessage
		var messageID, content, replyTo sql.NullString
		var messageCreatedAt sql.NullTime
		var senderID, senderName, senderUsername sql.NullString

		err := rows.Scan(
			&chat.ID, &chat.Type, &chat.Name, &chat.Description, &chat.UpdatedAt,
			&messageID, &content, &replyTo, &messageCreatedAt,
			&senderID, &senderName, &senderUsername,
			&chat.MemberCount,
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
				Sender: &user_readmodels.Sender{
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

	if direction == "newer" {
		for i, j := 0, len(chats)-1; i < j; i, j = i+1, j-1 {
			chats[i], chats[j] = chats[j], chats[i]
		}
	}

	hasNext, hasPrev, err := q.checkPaginationBounds(ctx, userID, chats)
	if err != nil {
		return nil, err
	}

	return &readmodels.ChatList{
		Chats:   chats,
		HasNext: hasNext,
		HasPrev: hasPrev,
	}, nil
}

func (q *ChatListQuery) getChatUpdatedAt(ctx context.Context, chatID string) (string, error) {
	query := `SELECT updated_at FROM chats WHERE id = $1`

	var updatedAt string
	err := q.db.QueryRowContext(ctx, query, chatID).Scan(&updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return updatedAt, nil
}

func (q *ChatListQuery) checkPaginationBounds(
	ctx context.Context,
	userID string,
	chats []*readmodels.ChatWithLastMessage,
) (bool, bool, error) {

	if len(chats) == 0 {
		return false, false, nil
	}

	var hasNext, hasPrev bool
	var err error

	firstChat := chats[0]
	lastChat := chats[len(chats)-1]

	hasNextQuery := `
        SELECT EXISTS(
            SELECT 1 FROM chats c
            JOIN members m ON c.id = m.chat_id
            WHERE m.user_id = $1 
            AND (c.updated_at > $2 OR (c.updated_at = $2 AND c.id > $3))
        )`
	err = q.db.QueryRowContext(ctx, hasNextQuery, userID, firstChat.UpdatedAt, firstChat.ID).Scan(&hasNext)
	if err != nil {
		return false, false, err
	}

	hasPrevQuery := `
        SELECT EXISTS(
            SELECT 1 FROM chats c
            JOIN members m ON c.id = m.chat_id
            WHERE m.user_id = $1 
            AND (c.updated_at < $2 OR (c.updated_at = $2 AND c.id < $3))
        )`
	err = q.db.QueryRowContext(ctx, hasPrevQuery, userID, lastChat.UpdatedAt, lastChat.ID).Scan(&hasPrev)
	if err != nil {
		return false, false, err
	}

	return hasNext, hasPrev, nil
}
