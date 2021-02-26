package errors

import (
	"fmt"
	"net/http"
)

type RestError struct {
	Err    string `json:"err"`
	Status int    `json:"status "`
	Msg    string `json:"msg"`
}

//TODO
func (r *RestError) Error() string {
	return "error"
}

func NewBadRequestError(msg string) *RestError {
	return &RestError{
		Err:    "bad_request",
		Status: http.StatusBadRequest,
		Msg:    msg,
	}
}

func NewBadRequestErrorf(format string, v ...interface{}) *RestError {
	return &RestError{
		Err:    "bad_request",
		Status: http.StatusBadRequest,
		Msg:    fmt.Sprintf(format, v...),
	}
}

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Err:    "internal_server_error",
		Status: http.StatusInternalServerError,
		Msg:    msg,
	}
}

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Err:    "not_found",
		Status: http.StatusNotFound,
		Msg:    msg,
	}
}

func NewNotFoundErrorf(format string, v ...interface{}) *RestError {
	return &RestError{
		Err:    "not_found",
		Status: http.StatusNotFound,
		Msg:    fmt.Sprintf(format, v...),
	}
}

func NewNotImplementingYet(msg string) *RestError {
	return &RestError{
		Err:    "not_implementing_yet",
		Status: http.StatusNotImplemented,
		Msg:    msg,
	}
}
