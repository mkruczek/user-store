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
		Err:    "bad request",
		Status: http.StatusBadRequest,
		Msg:    msg,
	}
}

func NewBadRequestErrorf(format string, v ...interface{}) *RestError {
	return &RestError{
		Err:    "bad request",
		Status: http.StatusBadRequest,
		Msg:    fmt.Sprintf(format, v...),
	}
}

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Err:    "internal server error",
		Status: http.StatusInternalServerError,
		Msg:    msg,
	}
}

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Err:    "not found",
		Status: http.StatusNotFound,
		Msg:    msg,
	}
}

func NewNotImplementingYet(msg string) *RestError {
	return &RestError{
		Err:    "not implementing yet",
		Status: http.StatusNotImplemented,
		Msg:    msg,
	}
}
