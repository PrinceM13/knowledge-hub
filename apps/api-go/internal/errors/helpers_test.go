package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRequest(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		expectCodeMatch bool
		expectMessage   string
	}{
		{
			name:            "generic validation error",
			err:             errors.New("invalid input"),
			expectCodeMatch: true,
			expectMessage:   "Invalid request payload",
		},
		{
			name:            "nil error",
			err:             nil,
			expectCodeMatch: true,
			expectMessage:   "Invalid request payload",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			appErr := InvalidRequest(tt.err)

			// assert
			assert.NotNil(t, appErr)
			if tt.expectCodeMatch {
				assert.Equal(t, ErrInvalidInput.Code, appErr.Code)
			}
			assert.Contains(t, appErr.Message, tt.expectMessage)
			assert.Equal(t, ErrInvalidInput.HTTPStatus, appErr.HTTPStatus)
		})
	}
}

func TestInvalidPathParam(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		paramName string
	}{
		{
			name:      "invalid integer param",
			err:       errors.New("strconv.ParseInt: parsing \"abc\": invalid syntax"),
			paramName: "id",
		},
		{
			name:      "invalid UUID param",
			err:       errors.New("invalid UUID format"),
			paramName: "userId",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			appErr := InvalidPathParam(tt.err, tt.paramName)

			// assert
			assert.NotNil(t, appErr)
			assert.Equal(t, ErrBadRequest.Code, appErr.Code)
			assert.Contains(t, appErr.Message, tt.paramName)
			assert.Contains(t, appErr.Message, "path parameter")
			assert.Equal(t, http.StatusBadRequest, appErr.HTTPStatus)
			assert.Equal(t, tt.err, appErr.Err)
		})
	}
}

func TestInvalidQueryParam(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		paramName string
	}{
		{
			name:      "invalid limit param",
			err:       errors.New("invalid integer"),
			paramName: "limit",
		},
		{
			name:      "invalid offset param",
			err:       errors.New("negative value"),
			paramName: "offset",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			appErr := InvalidQueryParam(tt.err, tt.paramName)

			// assert
			assert.NotNil(t, appErr)
			assert.Equal(t, ErrBadRequest.Code, appErr.Code)
			assert.Contains(t, appErr.Message, tt.paramName)
			assert.Contains(t, appErr.Message, "query parameter")
			assert.Equal(t, http.StatusBadRequest, appErr.HTTPStatus)
			assert.Equal(t, tt.err, appErr.Err)
		})
	}
}

func TestMissingField(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
	}{
		{
			name:      "missing email field",
			fieldName: "email",
		},
		{
			name:      "missing password field",
			fieldName: "password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			appErr := MissingField(tt.fieldName)

			// assert
			assert.NotNil(t, appErr)
			assert.Equal(t, ErrMissingField.Code, appErr.Code)
			assert.Contains(t, appErr.Message, tt.fieldName)
			assert.Contains(t, appErr.Message, "Required field")
			assert.Equal(t, ErrMissingField.HTTPStatus, appErr.HTTPStatus)
			assert.Nil(t, appErr.Err)
		})
	}
}

func TestMissingFields(t *testing.T) {
	tests := []struct {
		name       string
		fieldNames []string
	}{
		{
			name:       "missing single field",
			fieldNames: []string{"email"},
		},
		{
			name:       "missing multiple fields",
			fieldNames: []string{"email", "password", "name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			appErr := MissingFields(tt.fieldNames...)

			// assert
			assert.NotNil(t, appErr)
			assert.Equal(t, ErrMissingField.Code, appErr.Code)
			assert.Contains(t, appErr.Message, "Required fields")
			for _, field := range tt.fieldNames {
				assert.Contains(t, appErr.Message, field)
			}
			assert.Equal(t, ErrMissingField.HTTPStatus, appErr.HTTPStatus)
			assert.Nil(t, appErr.Err)
		})
	}
}
