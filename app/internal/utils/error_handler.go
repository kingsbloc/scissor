package utils

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Err            error  `json:"-"`               // low-level runtime error
	HTTPStatusCode int    `json:"-"`               // http response status code
	StatusText     string `json:"status"`          // user-level status message
	AppCode        int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText      string `json:"error,omitempty"` // application-level error message, for debugging
}

type ValidationError struct {
	Namespace string      `json:"namespace"`
	Field     string      `json:"field"`
	Type      string      `json:"type"`
	Tag       string      `json:"tag"`
	Value     interface{} `json:"value"`
}

func (e *ValidationError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusUnprocessableEntity)
	return nil
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ApiResponse{
		Status:  400,
		Message: err.Error(),
		Success: false,
	}
}

func ErrValidationRequest(err error, message string) render.Renderer {
	return &ApiResponse{
		Status:  http.StatusUnprocessableEntity,
		Message: message,
		Success: false,
		Data: map[string]interface{}{
			"errors": FormatBodyError(err),
		},
	}
}

func FormatBodyError(validationErr error) []ValidationError {
	var errors []ValidationError
	for _, err := range validationErr.(validator.ValidationErrors) {
		error := &ValidationError{
			Namespace: err.Namespace(),
			Field:     err.Field(),
			Type:      err.Type().Name(),
			Tag:       err.Tag(),
			Value:     err.Value(),
		}
		errors = append(errors, *error)
	}
	return errors
}
