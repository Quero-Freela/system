package utils

import (
	"encoding/json"
	"github.com/Quero-Freela/system/server/exceptions"
	"net/http"
)

type HttpError struct {
	Code   int
	Errors []string
}

type ResponseMessage struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []HttpError `json:"errors,omitempty"`
}

func NewResponseMessageError(err error, code int, data ...interface{}) *ResponseMessage {
	if err == nil {
		return &ResponseMessage{
			Data: data,
		}
	}

	errs := make([]string, 0)

	var er any = err
	if e, ok := er.(interface{ Unwrap() []error }); ok {
		for _, e := range e.Unwrap() {
			errs = append(errs, e.Error())
		}
	} else {
		errs = append(errs, err.Error())
	}

	return &ResponseMessage{
		Errors: []HttpError{
			{
				Code:   code,
				Errors: errs,
			},
		},
		Data: data,
	}
}

func WriteResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if _, ok := data.(ResponseMessage); !ok {
			data = ResponseMessage{
				Data: data,
			}
		}
	}

	jsonString, _ := json.Marshal(data)
	_, _ = w.Write(jsonString)
}

func WriteError(w http.ResponseWriter, err error, code int, status int, data ...interface{}) {
	d := NewResponseMessageError(err, code, data...)
	WriteResponse(w, d, status)
}

func Ok(w http.ResponseWriter, data interface{}) {
	WriteResponse(w, data, http.StatusOK)
}

func Error(w http.ResponseWriter, err error, code int, data ...interface{}) {
	switch {
	case exceptions.IsNfe(err):
		NotFound(w, err, code, data...)

	case exceptions.IsBre(err):
		BadRequest(w, err, code, data...)

	case exceptions.IsForbidden(err):
		Forbidden(w, err, code, data...)

	case exceptions.IsConflict(err):
		Conflict(w, err, code, data...)

	case exceptions.IsUnprocessableEntity(err):
		UnprocessableEntity(w, err, code, data...)

	case exceptions.IsTooManyRequests(err):
		TooManyRequests(w, err, code, data...)

	case exceptions.IsUnauthorized(err):
		Unauthorized(w, err, code, data...)

	default:
		InternalServerError(w, err, code, data...)
	}
}

func Success(w http.ResponseWriter, r *http.Request, data interface{}) {
	if data == nil {
		NoContent(w)
		return
	}

	if r.Method == http.MethodPost {
		Created(w, data)
		return
	}

	if r.Method == http.MethodPatch {
		PartialContent(w, data)
		return
	}

	if r.Method == http.MethodDelete {
		Accepted(w, data)
		return
	}

	if r.Method == http.MethodPut {
		ResetContent(w, data)
		return
	}

	Ok(w, data)
}

func BadRequest(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusBadRequest, data)
}

func InternalServerError(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusInternalServerError, data...)
}

func Unauthorized(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusUnauthorized, data...)
}

func Forbidden(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusForbidden, data...)
}

func NotFound(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusNotFound, data...)
}

func Conflict(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusConflict, data...)
}

func UnprocessableEntity(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusUnprocessableEntity, data...)
}

func TooManyRequests(w http.ResponseWriter, err error, code int, data ...interface{}) {
	WriteError(w, err, code, http.StatusTooManyRequests, data...)
}

func Created(w http.ResponseWriter, data interface{}) {
	WriteResponse(w, data, http.StatusCreated)
}

func NoContent(w http.ResponseWriter) {
	WriteResponse(w, nil, http.StatusNoContent)
}

func Accepted(w http.ResponseWriter, data interface{}) {
	WriteResponse(w, data, http.StatusAccepted)
}

func PartialContent(w http.ResponseWriter, data interface{}) {
	WriteResponse(w, data, http.StatusPartialContent)
}

func ResetContent(w http.ResponseWriter, data interface{}) {
	WriteResponse(w, data, http.StatusResetContent)
}
