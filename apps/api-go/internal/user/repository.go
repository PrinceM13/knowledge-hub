package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	List(ctx context.Context, limit, offset int) ([]*User, error)
}
