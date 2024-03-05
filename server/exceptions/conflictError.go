package exceptions

import (
	"errors"
	"fmt"
)

type ConflictError struct {
	Message string
	Code    int
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("Conflict [%d]: %s", e.Code, e.Message)
}

func IsConflict(err error) bool {
	var fe *ConflictError
	return errors.As(err, &fe)
}

func NewConflictErrorWithCode(message string, code int) *ConflictError {
	return &ConflictError{
		Message: message,
		Code:    code,
	}
}

func NewConflictWithCodef(code int, format string, a ...interface{}) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
		Code:    403,
	}
}

func NewConflictErrorf(format string, a ...interface{}) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf(format, a...),
		Code:    403,
	}
}
