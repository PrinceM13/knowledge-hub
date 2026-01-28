package errors

import "net/http"

// common application errors
var (
	// generic errors
	ErrInternal     = New("INTERNAL_ERROR", "An internal error occurred", http.StatusInternalServerError)
	ErrBadRequest   = New("BAD_REQUEST", "Invalid request", http.StatusBadRequest)
	ErrUnauthorized = New("UNAUTHORIZED", "Unauthorized", http.StatusUnauthorized)
	ErrForbidden    = New("FORBIDDEN", "Forbidden", http.StatusForbidden)
	ErrNotFound     = New("NOT_FOUND", "Resource not found", http.StatusNotFound)
	ErrConflict     = New("CONFLICT", "Resource already exists", http.StatusConflict)

	// validation errors
	ErrInvalidInput = New("INVALID_INPUT", "Input validation failed", http.StatusBadRequest)
	ErrMissingField = New("MISSING_FIELD", "Required field is missing", http.StatusBadRequest)
)
