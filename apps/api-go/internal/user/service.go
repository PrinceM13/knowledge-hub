package user

import (
	"context"
	"database/sql"
	"fmt"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, email, name string) (*User, error) {
	u := &User{
		Email: email,
		Name:  name,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) FindByID(ctx context.Context, id int64) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ListUsers(ctx context.Context, limit, offset int) ([]*User, error) {
	users, err := s.repo.List(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) RegisterUser(ctx context.Context, email, name string) (*User, error) {
	existingUser, err := s.repo.FindByEmail(ctx, email)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	return s.Create(ctx, email, name)
}
