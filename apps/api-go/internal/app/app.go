package app

import (
	"context"

	"github.com/PrinceM13/knowledge-hub-api/internal/user"
)

type App struct {
	User *user.Service
}

func New(userSvc *user.Service) *App {
	return &App{User: userSvc}
}

func (a *App) RegisterUser(ctx context.Context, email, name string) (*user.User, error) {
	return a.User.RegisterUser(ctx, email, name)
}

func (a *App) ListUsers(ctx context.Context, limit, offset int) ([]*user.User, error) {
	return a.User.ListUsers(ctx, limit, offset)
}

func (a *App) GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	return a.User.FindByID(ctx, id)
}
