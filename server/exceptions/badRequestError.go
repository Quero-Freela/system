package exceptions

import (
	"errors"
	"fmt"
)

type BadRequestError struct {
	Message string
	Code    int
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest [%d]: %s", e.Code, e.Message)
}

func IsBre(err error) bool {
	var be *BadRequestError
	return errors.As(err, &be)
}

func NewBadRequestErrorWithCode(message string, code int) *BadRequestError {
	return &BadRequestError{
		Message: message,
		Code:    code,
	}
}

func NewBadRequestWithCodef(code int, format string, a ...interface{}) *BadRequestError {
	return &BadRequestError{
		Message: fmt.Sprintf(format, a...),
		Code:    code,
	}
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
		Code:    404,
	}
}

func NewBadRequestErrorf(format string, a ...interface{}) *BadRequestError {
	return &BadRequestError{
		Message: fmt.Sprintf(format, a...),
		Code:    404,
	}
}
