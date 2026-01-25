package userdb

import (
	"context"
	"database/sql"

	"github.com/PrinceM13/knowledge-hub-api/internal/user"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) user.Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u *user.User) error {
	query := `
        INSERT INTO users (email, name)
        VALUES ($1, $2)
        RETURNING id, created_at
    `

	return r.db.QueryRowContext(ctx, query, u.Email, u.Name).Scan(&u.ID, &u.CreatedAt)
}

func (r *PostgresRepository) FindByID(ctx context.Context, id int64) (*user.User, error) {
	query := `
        SELECT id, email, name, created_at
        FROM users
        WHERE id = $1
    `

	u := &user.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT id, email, name, created_at
		FROM users
		WHERE email = $1
	`

	u := &user.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresRepository) List(ctx context.Context, limit, offset int) ([]*user.User, error) {
	query := `
		SELECT id, email, name, created_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User

	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
