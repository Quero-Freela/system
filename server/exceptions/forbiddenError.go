package exceptions

import (
	"errors"
	"fmt"
)

type ForbiddenError struct {
	Message string
	Code    int
}

func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("Forbidden [%d]: %s", e.Code, e.Message)
}

func IsForbidden(err error) bool {
	var fe *ForbiddenError
	return errors.As(err, &fe)
}

func NewForbiddenErrorWithCode(message string, code int) *ForbiddenError {
	return &ForbiddenError{
		Message: message,
		Code:    code,
	}
}

func NewForbiddenWithCodef(code int, format string, a ...interface{}) *ForbiddenError {
	return &ForbiddenError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{
		Message: message,
		Code:    403,
	}
}

func NewForbiddenErrorf(format string, a ...interface{}) *ForbiddenError {
	return &ForbiddenError{
		Message: fmt.Sprintf(format, a...),
		Code:    403,
	}
}
