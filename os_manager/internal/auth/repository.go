package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(
	ctx context.Context,
	name, email, password, role string,
) error {
	query := `
		INSERT INTO users (name, email, password, role)
		VALUES($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, name, email, password, role)
	return err
}
