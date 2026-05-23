package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/portfolio/api/internal/models"
)

// UserRepository handles database operations for ***REMOVED*** users.
type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// FindByEmail retrieves a user by email address.
// Returns the user with password hash for authentication purposes.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(ctx,
		`SELECT id, email, password_hash, name, created_at, updated_at
		FROM ***REMOVED***_users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

// FindByID retrieves a user by ID.
func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(ctx,
		`SELECT id, email, password_hash, name, created_at, updated_at
		FROM ***REMOVED***_users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

// UpdatePassword updates the password hash for a user.
func (r *UserRepository) UpdatePassword(ctx context.Context, id, passwordHash string) error {
	_, err := r.db.Exec(ctx,
		`UPDATE ***REMOVED***_users SET password_hash = $1, updated_at = $2 WHERE id = $3`,
		passwordHash, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}
	return nil
}

// UpdateProfile updates the name and email for a user.
func (r *UserRepository) UpdateProfile(ctx context.Context, id, name, email string) error {
	_, err := r.db.Exec(ctx,
		`UPDATE ***REMOVED***_users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`,
		name, email, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}
	return nil
}
