package user

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/PrinceM13/knowledge-hub-api/internal/errors"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, email, name string) (*User, error) {
	// validate email format, already validated at http layer but keeping for domain consistency
	if !emailRegex.MatchString(email) {
		return nil, ErrInvalidEmail
	}

	// validate name length
	if len(name) < 2 || len(name) > 100 {
		return nil, ErrInvalidName
	}

	u := &User{
		Email: email,
		Name:  name,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, errors.Wrap(err, errors.ErrInternal.Code, "Failed to create user", errors.ErrInternal.HTTPStatus)
	}

	return u, nil
}

func (s *Service) FindByID(ctx context.Context, id int64) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.Wrap(err, errors.ErrInternal.Code, "Failed to fetch user", errors.ErrInternal.HTTPStatus)
	}

	return user, nil
}

func (s *Service) ListUsers(ctx context.Context, limit, offset int) ([]*User, error) {
	users, err := s.repo.List(ctx, limit, offset)

	if err != nil {
		return nil, errors.Wrap(err, errors.ErrInternal.Code, "Failed to fetch users", errors.ErrInternal.HTTPStatus)
	}

	return users, nil
}

func (s *Service) RegisterUser(ctx context.Context, email, name string) (*User, error) {
	existingUser, err := s.repo.FindByEmail(ctx, email)

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, errors.ErrInternal.Code, "Failed to check existing user", errors.ErrInternal.HTTPStatus)
	}

	if existingUser != nil {
		return nil, ErrDuplicateEmail
	}

	return s.Create(ctx, email, name)
}
