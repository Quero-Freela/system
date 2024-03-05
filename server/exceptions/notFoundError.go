package exceptions

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Message string
	Code    int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound [%d]: %s", e.Code, e.Message)
}

func IsNfe(err error) bool {
	var notFoundError *NotFoundError
	return errors.As(err, &notFoundError)
}

func NewNotFoundErrorWithCode(message string, code int) *NotFoundError {
	return &NotFoundError{
		Message: message,
		Code:    code,
	}
}

func NewNotFoundWithCodef(code int, format string, a ...interface{}) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
		Code:    404,
	}
}

func NewNotFoundErrorf(format string, a ...interface{}) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf(format, a...),
		Code:    404,
	}
}
