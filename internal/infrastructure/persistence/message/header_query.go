package persistence

import (
	"context"
	"database/sql"
	"main/internal/application/message/queries"
	"main/internal/application/message/readmodels"
	"main/pkg"
)

type HeaderQuery struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewHeaderQuery(db pkg.PostgresDB, logger pkg.Logger) queries.HeaderQuery {
	return &HeaderQuery{
		db:     db,
		logger: logger,
	}
}

func (q *HeaderQuery) Query(chatID string) (*readmodels.ChatHeader, error) {
	ctx := context.Background()

	query := `
		SELECT 
			id, 
			type, 
			name
		FROM chats 
		WHERE id = $1
	`

	var header readmodels.ChatHeader
	err := q.db.QueryRowContext(ctx, query, chatID).Scan(
		&header.ID,
		&header.Type,
		&header.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		q.logger.Error("Failed to query chat header", "chatID", chatID, "error", err)
		return nil, err
	}

	return &header, nil
}
