package error

import (
	"fmt"
	"net/http"
)

type Error struct {
	code     int      `json:"code,omitempty"`
	msg      string   `json:"msg,omitempty"`
	details  []string `json:"details,omitempty"`
}

var _codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("error code %d exists, replace with ohters", code))
	}

	_codes[code] = msg

	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetail(details ...string) *Error {
	newError := *e
	newError.details = make([]string, 0)

	for _, detail := range details {
		newError.details = append(newError.details, detail)
	}

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case InvalidParams.Code():
		return http.StatusBadRequest
	case ServerError.Code():
		return http.StatusInternalServerError
	case AuthorizedExpire.Code():
		fallthrough
	case UnauthorizedRights.Code():
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

