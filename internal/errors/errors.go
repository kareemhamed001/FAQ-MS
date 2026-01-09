package errors

import "net/http"

// AppError represents a typed application error with HTTP semantics.
type AppError struct {
	Code    string
	Message string
	Status  int
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

// Predefined common errors.
var (
	ErrUnauthorized = &AppError{Code: "UNAUTHORIZED", Message: "access denied", Status: http.StatusUnauthorized}
	ErrForbidden    = &AppError{Code: "FORBIDDEN", Message: "forbidden", Status: http.StatusForbidden}
	ErrNotFound     = &AppError{Code: "NOT_FOUND", Message: "resource not found", Status: http.StatusNotFound}
	ErrInvalidInput = &AppError{Code: "INVALID_INPUT", Message: "invalid input", Status: http.StatusBadRequest}
	ErrEmailExists  = &AppError{Code: "EMAIL_EXISTS", Message: "email already registered", Status: http.StatusConflict}
	ErrDuplicate    = &AppError{Code: "DUPLICATE", Message: "duplicate resource", Status: http.StatusConflict}
	ErrInternal     = &AppError{Code: "INTERNAL_ERROR", Message: "internal error", Status: http.StatusInternalServerError}
)

// Wrap preserves an inner error while keeping the envelope metadata.
func Wrap(app *AppError, err error) *AppError {
	if err == nil {
		return app
	}
	return &AppError{
		Code:    app.Code,
		Message: app.Message,
		Status:  app.Status,
		Err:     err,
	}
}
