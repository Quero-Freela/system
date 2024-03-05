package exceptions

import (
	"errors"
	"fmt"
)

type UnprocessableEntityError struct {
	Message string
	Code    int
}

func (e *UnprocessableEntityError) Error() string {
	return fmt.Sprintf("UnprocessableEntity [%d]: %s", e.Code, e.Message)
}

func IsUnprocessableEntity(err error) bool {
	var fe *UnprocessableEntityError
	return errors.As(err, &fe)
}

func NewUnprocessableEntityErrorWithCode(message string, code int) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		Message: message,
		Code:    code,
	}
}

func NewUnprocessableEntityWithCodef(code int, format string, a ...interface{}) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		Message: message,
		Code:    403,
	}
}

func NewUnprocessableEntityErrorf(format string, a ...interface{}) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		Message: fmt.Sprintf(format, a...),
		Code:    403,
	}
}
