package errors

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code       string // Machine-readable error code
	Message    string // Human-readable error message
	HTTPStatus int    // HTTP status code to return
	Err        error  // Underlying error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.Err)
	}
	return e.Message
}

// convert to underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// create a new fresh AppError
func New(code, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

// convert an existing error into an AppError
func Wrap(err error, code, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		Err:        err,
	}
}

// inspect an error and see if it's an AppError
func AsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}
