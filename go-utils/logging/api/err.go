package api

import (
	"fmt"
	"net/http"
	"runtime"
)

const (
	ErrBadRequest     = 0
	ErrInternalServer = 1
)

var errorMap = map[int]Error{
	ErrInternalServer: {"internal", "server_error", "internal server error", "", nil, ErrInternalServer, http.StatusInternalServerError},
}

type Error struct {
	Code        string                 `json:"code"`
	Type        string                 `json:"type"`
	Message     string                 `json:"message"`
	Description string                 `json:"description"`
	Params      map[string]interface{} `json:"-"`
	ErrCode     int                    `json:"-"`
	HttpStatus  int                    `json:"-"`
}

func (e Error) Error() string {
	if e.Params == nil {
		return fmt.Sprintf("%d: %s (%s)", e.ErrCode, e.Code, e.Message)
	}
	return fmt.Sprintf("%d: %s (%s) params: %s", e.ErrCode, e.Code, e.Message, e.Params)
}

func NewError(code int) error {
	_, fn, line, _ := runtime.Caller(1)

	e, ok := errorMap[code]
	if !ok {
		e = errorMap[ErrBadRequest] // 0
	}

	e.Params = map[string]interface{}{
		"where": fmt.Sprintf("%s:%d", fn, line),
	}

	return e
}

func NewCustomError(code int, desc string, params map[string]interface{}) error {
	e, ok := errorMap[code]
	if !ok {
		e = errorMap[ErrBadRequest] // 0
	}

	if desc != "" {
		e.Message = desc
	}

	e.Description = desc

	var (
		fn   string
		line int
	)

	if code == ErrInternalServer {
		_, fn, line, _ = runtime.Caller(2)
	} else {
		_, fn, line, _ = runtime.Caller(1)
	}

	if params == nil {
		e.Params = map[string]interface{}{
			"where": fmt.Sprintf("%s:%d", fn, line),
		}
	} else {
		e.Params = params
		e.Params["where"] = fmt.Sprintf("%s:%d", fn, line)
	}

	return e
}

func NewInternalError(desc string, params map[string]interface{}) error {
	return NewCustomError(ErrInternalServer, desc, params)
}
