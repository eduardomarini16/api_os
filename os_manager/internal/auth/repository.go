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

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (
	id int64,
	hash string,
	role string,
	err error,
) {
	query := `
		SELECT id, password, role
		FROM users
		WHERE email = $1
	`

	err = r.db.QueryRow(ctx, query).Scan(&id, &hash, &role)
	return
}
