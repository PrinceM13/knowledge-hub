package user

import (
	"github.com/PrinceM13/knowledge-hub-api/internal/errors"
)

// domain-specific errors that need custom codes/messages
var (
	ErrDuplicateEmail = errors.New("DUPLICATE_EMAIL", "Email already exists", errors.ErrConflict.HTTPStatus)
	ErrInvalidName    = errors.New("INVALID_NAME", "Name must be between 2 and 100 characters", errors.ErrBadRequest.HTTPStatus)

	// validate email format, already validated in service at http layer but keeping for domain consistency
	ErrInvalidEmail = errors.New("INVALID_EMAIL", "Invalid email format", errors.ErrBadRequest.HTTPStatus)
)
