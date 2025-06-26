package repositories

import (
	"context"
	"database/sql"
	"errors"

	"mfo-service/internal/apperrs"
	"mfo-service/internal/domain"
)

type ColdUsersRepository struct {
	db *sql.DB
}

func NewColdUsersRepository(db *sql.DB) *ColdUsersRepository {
	return &ColdUsersRepository{db: db}
}

func (r *ColdUsersRepository) FindByPhone(ctx context.Context, phone int) (*domain.ColdUsers, error) {
	query := `SELECT phone, credits FROM cold_users WHERE phone = $1`
	coldUsers := &domain.ColdUsers{}

	err := r.db.QueryRowContext(ctx, query, phone).Scan(&coldUsers.Phone, &coldUsers.Credits)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrs.ErrNotFound
		}
		return nil, err
	}

	return coldUsers, nil
}
