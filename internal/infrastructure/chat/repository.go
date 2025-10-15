package chat_infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	chat_domain "main/internal/domain/chat"
	"main/pkg"
)

type ChatRepository struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewChatRepository(db pkg.PostgresDB, logger pkg.Logger) chat_domain.ChatRepository {
	if err := createChatTables(db); err != nil {
		logger.Error("failed to create chat tables", err)
	}

	return &ChatRepository{
		db:     db,
		logger: logger,
	}
}

func createChatTables(db pkg.PostgresDB) error {
	query := `
		CREATE TABLE IF NOT EXISTS chats (
			id VARCHAR PRIMARY KEY,
			type VARCHAR NOT NULL CHECK (type IN ('direct', 'group')),
			name VARCHAR,
			description TEXT DEFAULT '',
			created_by VARCHAR NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);

		CREATE TABLE IF NOT EXISTS members (
			chat_id VARCHAR REFERENCES chats(id) ON DELETE CASCADE,
			user_id VARCHAR NOT NULL,
			role VARCHAR DEFAULT 'member',
			joined_at TIMESTAMP NOT NULL,
			PRIMARY KEY (chat_id, user_id)
		);

		CREATE INDEX IF NOT EXISTS idx_chats_type ON chats(type);
		CREATE INDEX IF NOT EXISTS idx_chats_updated ON chats(updated_at DESC);
		CREATE INDEX IF NOT EXISTS idx_members_user_id ON members(user_id);
		CREATE INDEX IF NOT EXISTS idx_members_chat_id ON members(chat_id);
	`

	_, err := db.Exec(query)
	return err
}

func (r *ChatRepository) Create(chat *chat_domain.Chat) error {
	if chat.ID == "" {
		chat.ID = generateID()
	}
	now := time.Now()
	chat.CreatedAt = now
	chat.UpdatedAt = now

	return r.db.WithTransaction(context.Background(), func(tx *sql.Tx) error {
		chatQuery := `
			INSERT INTO chats (id, type, name, description, created_by, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
		_, err := tx.Exec(
			chatQuery,
			chat.ID,
			chat.Type,
			chat.Name,
			chat.Description,
			chat.CreatedBy,
			chat.CreatedAt,
			chat.UpdatedAt,
		)
		if err != nil {
			return err
		}

		memberQuery := `
			INSERT INTO members (chat_id, user_id, role, joined_at)
			VALUES ($1, $2, $3, $4)
		`
		for _, member := range chat.Members {
			_, err := tx.Exec(
				memberQuery,
				chat.ID,
				member.UserID,
				member.Role,
				now,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *ChatRepository) GetByID(id string) (*chat_domain.Chat, error) {
	chatQuery := `
		SELECT id, type, name, description, created_by, created_at, updated_at
		FROM chats 
		WHERE id = $1
	`

	var chat chat_domain.Chat
	err := r.db.QueryRow(chatQuery, id).Scan(
		&chat.ID,
		&chat.Type,
		&chat.Name,
		&chat.Description,
		&chat.CreatedBy,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	membersQuery := `
		SELECT user_id, role, joined_at
		FROM members 
		WHERE chat_id = $1
		ORDER BY joined_at
	`

	rows, err := r.db.Query(membersQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []chat_domain.ChatMember
	for rows.Next() {
		var member chat_domain.ChatMember
		err := rows.Scan(
			&member.UserID,
			&member.Role,
			&member.JoinedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	chat.Members = members
	return &chat, nil
}

func (r *ChatRepository) GetMemberChats(userID string, offset, limit int) (*chat_domain.ChatList, error) {
	ctx := context.Background()

	var total int
	countQuery := `
		SELECT COUNT(DISTINCT c.id)
		FROM chats c
		JOIN members m ON c.id = m.chat_id
		WHERE m.user_id = $1
	`
	err := r.db.QueryRowContext(ctx, countQuery, userID).Scan(&total)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT c.id, c.type, c.name, c.description, c.created_by, c.created_at, c.updated_at
		FROM chats c
		JOIN members m ON c.id = m.chat_id
		WHERE m.user_id = $1
		GROUP BY c.id
		ORDER BY c.updated_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*chat_domain.Chat
	for rows.Next() {
		var chat chat_domain.Chat
		err := rows.Scan(
			&chat.ID,
			&chat.Type,
			&chat.Name,
			&chat.Description,
			&chat.CreatedBy,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		chats = append(chats, &chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &chat_domain.ChatList{
		Chats:  chats,
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}, nil
}

func (r *ChatRepository) Update(chat *chat_domain.Chat) error {
	chat.UpdatedAt = time.Now()

	query := `
		UPDATE chats 
		SET name = $1, description = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(
		query,
		chat.Name,
		chat.Description,
		chat.UpdatedAt,
		chat.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("chat not found: %s", chat.ID)
	}

	return nil
}

func (r *ChatRepository) Delete(id string) error {
	query := `DELETE FROM chats WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("chat not found: %s", id)
	}

	return nil
}

func (r *ChatRepository) FindMemberChat(userID, keyWord string, offset, limit int) (*chat_domain.ChatList, error) {
	ctx := context.Background()

	query := `
		SELECT DISTINCT c.id, c.type, c.name, c.description, c.created_by, c.created_at, c.updated_at
		FROM chats c
		JOIN members m ON c.id = m.chat_id
		WHERE m.user_id = $1
		AND (
			(c.type = 'group' AND c.name ILIKE $2)
			OR
			(c.type = 'direct' AND EXISTS (
				SELECT 1 FROM members m2
				JOIN users u ON m2.user_id = u.id
				WHERE m2.chat_id = c.id 
				AND m2.user_id != $1
				AND (u.username ILIKE $2 OR u.name ILIKE $2)
			))
		)
		ORDER BY c.updated_at DESC
		LIMIT $3 OFFSET $4
	`

	searchPattern := "%" + keyWord + "%"
	rows, err := r.db.QueryContext(ctx, query, userID, searchPattern, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*chat_domain.Chat
	for rows.Next() {
		var chat chat_domain.Chat
		err := rows.Scan(
			&chat.ID,
			&chat.Type,
			&chat.Name,
			&chat.Description,
			&chat.CreatedBy,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		chats = append(chats, &chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var total int
	countQuery := `
		SELECT COUNT(DISTINCT c.id)
		FROM chats c
		JOIN members m ON c.id = m.chat_id
		WHERE m.user_id = $1
		AND (
			(c.type = 'group' AND c.name ILIKE $2)
			OR
			(c.type = 'direct' AND EXISTS (
				SELECT 1 FROM members m2
				JOIN users u ON m2.user_id = u.id
				WHERE m2.chat_id = c.id 
				AND m2.user_id != $1
				AND (u.username ILIKE $2 OR u.name ILIKE $2)
			))
		)
	`
	err = r.db.QueryRowContext(ctx, countQuery, userID, searchPattern).Scan(&total)
	if err != nil {
		return nil, err
	}

	return &chat_domain.ChatList{
		Chats:  chats,
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}, nil
}

func (r *ChatRepository) GetDirectChat(firstUserID, secondUserID string) (*chat_domain.Chat, error) {
	query := `
		SELECT c.id, c.type, c.name, c.description, c.created_by, c.created_at, c.updated_at
		FROM chats c
		WHERE c.type = 'direct'
		AND EXISTS (
			SELECT 1 FROM members m 
			WHERE m.chat_id = c.id AND m.user_id = $1
		)
		AND EXISTS (
			SELECT 1 FROM members m 
			WHERE m.chat_id = c.id AND m.user_id = $2
		)
		AND (
			SELECT COUNT(*) FROM members m 
			WHERE m.chat_id = c.id
		) = 2
		LIMIT 1
	`

	var chat chat_domain.Chat
	err := r.db.QueryRow(query, firstUserID, secondUserID).Scan(
		&chat.ID,
		&chat.Type,
		&chat.Name,
		&chat.Description,
		&chat.CreatedBy,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	members, err := r.getChatMembers(chat.ID)
	if err != nil {
		return nil, err
	}
	chat.Members = members

	return &chat, nil
}

func (r *ChatRepository) MarkUpdated(chatID string, updateTime time.Time) error {
	query := `UPDATE chats SET updated_at = $1 WHERE id = $2`

	result, err := r.db.Exec(query, updateTime, chatID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("chat not found: %s", chatID)
	}

	return nil
}

func (r *ChatRepository) getChatMembers(chatID string) ([]chat_domain.ChatMember, error) {
	query := `
		SELECT user_id, role, joined_at
		FROM members 
		WHERE chat_id = $1
		ORDER BY joined_at
	`

	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []chat_domain.ChatMember
	for rows.Next() {
		var member chat_domain.ChatMember
		err := rows.Scan(
			&member.UserID,
			&member.Role,
			&member.JoinedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func generateID() string {
	return fmt.Sprintf("chat_%d", time.Now().UnixNano())
}
