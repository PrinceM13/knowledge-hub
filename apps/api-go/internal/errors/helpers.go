package errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// wraps request validation errors with detailed field information
func InvalidRequest(err error) *AppError {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		fieldErrors := make([]string, 0, len(validationErrors))
		for _, e := range validationErrors {
			fieldErrors = append(fieldErrors, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
		}
		message := fmt.Sprintf("Invalid request: %v", fieldErrors)
		return Wrap(err, ErrInvalidInput.Code, message, ErrInvalidInput.HTTPStatus)
	}

	return Wrap(err, ErrInvalidInput.Code, "Invalid request payload", ErrInvalidInput.HTTPStatus)
}

// for a specific parameter
func InvalidPathParam(err error, paramName string) *AppError {
	message := fmt.Sprintf("Invalid path parameter: %s", paramName)
	return Wrap(err, ErrBadRequest.Code, message, ErrBadRequest.HTTPStatus)
}

// for a specific query parameter
func InvalidQueryParam(err error, paramName string) *AppError {
	message := fmt.Sprintf("Invalid query parameter: %s", paramName)
	return Wrap(err, ErrBadRequest.Code, message, ErrBadRequest.HTTPStatus)
}

// for missing required field(s) when manual validation is needed
func MissingField(fieldName string) *AppError {
	message := fmt.Sprintf("Required field is missing: %s", fieldName)
	return Wrap(nil, ErrMissingField.Code, message, ErrMissingField.HTTPStatus)
}

// for missing multiple required fields
func MissingFields(fieldNames ...string) *AppError {
	message := fmt.Sprintf("Required fields are missing: %v", fieldNames)
	return Wrap(nil, ErrMissingField.Code, message, ErrMissingField.HTTPStatus)
}
