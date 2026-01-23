package user

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u *User) error {
	query := `
        INSERT INTO users (email, name)
        VALUES ($1, $2)
        RETURNING id, created_at
    `

	return r.db.QueryRowContext(ctx, query, u.Email, u.Name).Scan(&u.ID, &u.CreatedAt)
}

func (r *PostgresRepository) FindByID(ctx context.Context, id int64) (*User, error) {
	query := `
        SELECT id, email, name, created_at
        FROM users
        WHERE id = $1
    `

	u := &User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}
