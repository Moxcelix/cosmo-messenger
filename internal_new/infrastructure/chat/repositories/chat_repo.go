package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/repositories"
	"main/pkg"
	"time"
)

type PostgresChatRepository struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewPostgresChatRepository(db pkg.PostgresDB, logger pkg.Logger) repositories.ChatRepository {
	if err := createChatTables(db); err != nil {
		logger.Error("failed to create chat tables", err)
	}

	return &PostgresChatRepository{
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

func (r *PostgresChatRepository) Create(chat *models.Chat) error {
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
			member.JoinedAt = now
			_, err := tx.Exec(
				memberQuery,
				chat.ID,
				member.UserID,
				member.Role,
				member.JoinedAt,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *PostgresChatRepository) GetChatById(id string) (*models.Chat, error) {
	chatQuery := `
		SELECT id, type, name, description, created_by, created_at, updated_at
		FROM chats 
		WHERE id = $1
	`

	var chat models.Chat
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

	var members []*models.ChatMember
	for rows.Next() {
		var member models.ChatMember
		err := rows.Scan(
			&member.UserID,
			&member.Role,
			&member.JoinedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, &member)
	}

	chat.Members = members
	return &chat, nil
}

func (r *PostgresChatRepository) Update(chat *models.Chat) error {
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

func (r *PostgresChatRepository) Delete(id string) error {
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

func (r *PostgresChatRepository) GetDirectChat(firstUserID, secondUserID string) (*models.Chat, error) {
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

	var chat models.Chat
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

func (r *PostgresChatRepository) MarkUpdated(chatID string, updateTime time.Time) error {
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

func (r *PostgresChatRepository) getChatMembers(chatID string) ([]*models.ChatMember, error) {
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

	var members []*models.ChatMember
	for rows.Next() {
		member := &models.ChatMember{}
		err := rows.Scan(
			&member.UserID,
			&member.Role,
			&member.JoinedAt)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (r *PostgresChatRepository) ChatExists(chatId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM chats WHERE id = $1)`

	var exists bool
	err := r.db.QueryRow(query, chatId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresChatRepository) DirectChatExists(firstUserID, secondUserID string) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 FROM chats c
            INNER JOIN members m1 ON c.id = m1.chat_id AND m1.user_id = $1
            INNER JOIN members m2 ON c.id = m2.chat_id AND m2.user_id = $2
            WHERE c.type = 'direct'
        )
    `

	var exists bool
	err := r.db.QueryRow(query, firstUserID, secondUserID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresChatRepository) UserInChat(userId, chatId string) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 FROM members 
            WHERE chat_id = $1 AND user_id = $2
        )
    `

	var exists bool
	err := r.db.QueryRow(query, chatId, userId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func generateID() string {
	return fmt.Sprintf("chat_%d", time.Now().UnixNano())
}
