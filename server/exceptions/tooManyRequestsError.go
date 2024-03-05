package exceptions

import (
	"errors"
	"fmt"
)

type TooManyRequestsError struct {
	Message string
	Code    int
}

func (e *TooManyRequestsError) Error() string {
	return fmt.Sprintf("TooManyRequests [%d]: %s", e.Code, e.Message)
}

func IsTooManyRequests(err error) bool {
	var fe *TooManyRequestsError
	return errors.As(err, &fe)
}

func NewTooManyRequestsErrorWithCode(message string, code int) *TooManyRequestsError {
	return &TooManyRequestsError{
		Message: message,
		Code:    code,
	}
}

func NewTooManyRequestsWithCodef(code int, format string, a ...interface{}) *TooManyRequestsError {
	return &TooManyRequestsError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewTooManyRequestsError(message string) *TooManyRequestsError {
	return &TooManyRequestsError{
		Message: message,
		Code:    403,
	}
}

func NewTooManyRequestsErrorf(format string, a ...interface{}) *TooManyRequestsError {
	return &TooManyRequestsError{
		Message: fmt.Sprintf(format, a...),
		Code:    403,
	}
}
