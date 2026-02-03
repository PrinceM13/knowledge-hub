package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppError_Error(t *testing.T) {
	tests := []struct {
		name           string
		appError       *AppError
		expectedString string
	}{
		{
			name: "error with underlying error",
			appError: &AppError{
				Code:       "TEST_ERROR",
				Message:    "Failed to process",
				HTTPStatus: http.StatusInternalServerError,
				Err:        errors.New("connection timeout"),
			},
			expectedString: "Failed to process: connection timeout",
		},
		{
			name: "error without underlying error",
			appError: &AppError{
				Code:       "NOT_FOUND",
				Message:    "Resource not found",
				HTTPStatus: http.StatusNotFound,
				Err:        nil,
			},
			expectedString: "Resource not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			result := tt.appError.Error()

			// assert
			assert.Equal(t, tt.expectedString, result)
		})
	}
}

func TestAppError_Unwrap(t *testing.T) {
	tests := []struct {
		name          string
		appError      *AppError
		expectedError error
	}{
		{
			name: "unwrap with underlying error",
			appError: &AppError{
				Code:    "TEST",
				Message: "Test",
				Err:     errors.New("original error"),
			},
			expectedError: errors.New("original error"),
		},
		{
			name: "unwrap without underlying error",
			appError: &AppError{
				Code:    "TEST",
				Message: "Test",
				Err:     nil,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			result := tt.appError.Unwrap()

			// assert
			if tt.expectedError == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedError.Error(), result.Error())
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		code       string
		message    string
		httpStatus int
	}{
		{
			name:       "create new error",
			code:       "TEST_ERROR",
			message:    "This is a test error",
			httpStatus: http.StatusBadRequest,
		},
		{
			name:       "create internal error",
			code:       "INTERNAL_ERROR",
			message:    "Something went wrong",
			httpStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			err := New(tt.code, tt.message, tt.httpStatus)

			// assert
			assert.NotNil(t, err)
			assert.Equal(t, tt.code, err.Code)
			assert.Equal(t, tt.message, err.Message)
			assert.Equal(t, tt.httpStatus, err.HTTPStatus)
			assert.Nil(t, err.Err)
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name          string
		originalError error
		code          string
		message       string
		httpStatus    int
	}{
		{
			name:          "wrap existing error",
			originalError: errors.New("database connection failed"),
			code:          "DB_ERROR",
			message:       "Failed to connect to database",
			httpStatus:    http.StatusInternalServerError,
		},
		{
			name:          "wrap nil error",
			originalError: nil,
			code:          "VALIDATION_ERROR",
			message:       "Validation failed",
			httpStatus:    http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			err := Wrap(tt.originalError, tt.code, tt.message, tt.httpStatus)

			// assert
			assert.NotNil(t, err)
			assert.Equal(t, tt.code, err.Code)
			assert.Equal(t, tt.message, err.Message)
			assert.Equal(t, tt.httpStatus, err.HTTPStatus)
			assert.Equal(t, tt.originalError, err.Err)
		})
	}
}

func TestAsAppError(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		expectFound bool
		expectCode  string
	}{
		{
			name: "convert AppError successfully",
			err: &AppError{
				Code:       "TEST_ERROR",
				Message:    "Test message",
				HTTPStatus: http.StatusBadRequest,
			},
			expectFound: true,
			expectCode:  "TEST_ERROR",
		},
		{
			name: "wrapped AppError",
			err: Wrap(
				&AppError{
					Code:       "INNER_ERROR",
					Message:    "Inner message",
					HTTPStatus: http.StatusInternalServerError,
				},
				"OUTER_ERROR",
				"Outer message",
				http.StatusBadGateway,
			),
			expectFound: true,
			expectCode:  "OUTER_ERROR",
		},
		{
			name:        "standard error cannot be converted",
			err:         errors.New("standard error"),
			expectFound: false,
			expectCode:  "",
		},
		{
			name:        "nil error",
			err:         nil,
			expectFound: false,
			expectCode:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			appErr, found := AsAppError(tt.err)

			// assert
			assert.Equal(t, tt.expectFound, found)
			if tt.expectFound {
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.expectCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
			}
		})
	}
}

func TestPredefinedErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        *AppError
		expectCode string
		expectHTTP int
	}{
		{
			name:       "ErrInternal",
			err:        ErrInternal,
			expectCode: "INTERNAL_ERROR",
			expectHTTP: http.StatusInternalServerError,
		},
		{
			name:       "ErrBadRequest",
			err:        ErrBadRequest,
			expectCode: "BAD_REQUEST",
			expectHTTP: http.StatusBadRequest,
		},
		{
			name:       "ErrUnauthorized",
			err:        ErrUnauthorized,
			expectCode: "UNAUTHORIZED",
			expectHTTP: http.StatusUnauthorized,
		},
		{
			name:       "ErrForbidden",
			err:        ErrForbidden,
			expectCode: "FORBIDDEN",
			expectHTTP: http.StatusForbidden,
		},
		{
			name:       "ErrNotFound",
			err:        ErrNotFound,
			expectCode: "NOT_FOUND",
			expectHTTP: http.StatusNotFound,
		},
		{
			name:       "ErrConflict",
			err:        ErrConflict,
			expectCode: "CONFLICT",
			expectHTTP: http.StatusConflict,
		},
		{
			name:       "ErrInvalidInput",
			err:        ErrInvalidInput,
			expectCode: "INVALID_INPUT",
			expectHTTP: http.StatusBadRequest,
		},
		{
			name:       "ErrMissingField",
			err:        ErrMissingField,
			expectCode: "MISSING_FIELD",
			expectHTTP: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert
			assert.NotNil(t, tt.err)
			assert.Equal(t, tt.expectCode, tt.err.Code)
			assert.Equal(t, tt.expectHTTP, tt.err.HTTPStatus)
			assert.NotEmpty(t, tt.err.Message)
		})
	}
}
