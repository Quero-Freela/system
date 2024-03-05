package exceptions

import (
	"errors"
	"fmt"
)

type UnauthorizedError struct {
	Message string
	Code    int
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("Unauthorized [%d]: %s", e.Code, e.Message)
}

func IsUnauthorized(err error) bool {
	var fe *UnauthorizedError
	return errors.As(err, &fe)
}

func NewUnauthorizedErrorWithCode(message string, code int) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
		Code:    code,
	}
}

func NewUnauthorizedWithCodef(code int, format string, a ...interface{}) *UnauthorizedError {
	return &UnauthorizedError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
		Code:    403,
	}
}

func NewUnauthorizedErrorf(format string, a ...interface{}) *UnauthorizedError {
	return &UnauthorizedError{
		Message: fmt.Sprintf(format, a...),
		Code:    403,
	}
}
