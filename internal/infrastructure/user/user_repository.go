package user_infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	user_domain "main/internal/domain/user"
	"main/pkg"
)

type UserRepository struct {
	db     pkg.PostgresDB
	logger pkg.Logger
}

func NewUserRepository(db pkg.PostgresDB, logger pkg.Logger) user_domain.UserRepository {
	if err := createUserTable(db); err != nil {
		logger.Error("failed to create users table", err)
	}

	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func createUserTable(db pkg.PostgresDB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR PRIMARY KEY,
			name VARCHAR NOT NULL,
			username VARCHAR UNIQUE NOT NULL,
			password_hash VARCHAR NOT NULL,
			bio TEXT DEFAULT '',
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);

		CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username);
		CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at DESC);
	`

	_, err := db.Exec(query)
	return err
}

func (r *UserRepository) GetUserById(userId string) (*user_domain.User, error) {
	query := `
		SELECT id, name, username, password_hash, bio, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`

	var user user_domain.User
	err := r.db.QueryRow(query, userId).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.PasswordHash,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*user_domain.User, error) {
	query := `
		SELECT id, name, username, password_hash, bio, created_at, updated_at 
		FROM users 
		WHERE username = $1
	`

	var user user_domain.User
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.PasswordHash,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *user_domain.User) error {
	if user.ID == "" {
		user.ID = generateID()
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `
		INSERT INTO users (id, name, username, password_hash, bio, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Name,
		user.Username,
		user.PasswordHash,
		user.Bio,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		if isUniqueConstraintError(err) {
			return user_domain.ErrUsernameAlreadyTaken
		}
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUserByUsername(username string) error {
	query := `DELETE FROM users WHERE username = $1`

	result, err := r.db.Exec(query, username)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return user_domain.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) DeleteUserById(userId string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return user_domain.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *user_domain.User) error {
	user.UpdatedAt = time.Now()

	query := `
		UPDATE users 
		SET name = $1, username = $2, password_hash = $3, bio = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(
		query,
		user.Name,
		user.Username,
		user.PasswordHash,
		user.Bio,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		if isUniqueConstraintError(err) {
			return user_domain.ErrUsernameAlreadyTaken
		}
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return user_domain.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) GetUsersByRange(offset, limit int) (*user_domain.UsersList, error) {
	ctx := context.Background()

	var total int
	countQuery := `SELECT COUNT(*) FROM users`
	err := r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, name, username, password_hash, bio, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user_domain.User
	for rows.Next() {
		var user user_domain.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.PasswordHash,
			&user.Bio,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user_domain.UsersList{
		Users:  users,
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}, nil
}

func (r *UserRepository) UserExists(userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`

	var exists bool
	err := r.db.QueryRow(query, userId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func generateID() string {
	return fmt.Sprintf("user_%d", time.Now().UnixNano())
}

func isUniqueConstraintError(err error) bool {
	return err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\""
}
