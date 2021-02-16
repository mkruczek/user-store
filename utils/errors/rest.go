package errors

import "net/http"

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

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Err:    "internal server error",
		Status: http.StatusInternalServerError,
		Msg:    msg,
	}
}
