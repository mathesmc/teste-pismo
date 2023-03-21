package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrNotFound              = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
	ErrMismatchOperationType = fmt.Errorf("invalid_amount: operation type does not match with especific operation type")
)

func ErrInvalidDocument(s string) error {
	return fmt.Errorf("invalid_document_number: %s", s)

}

func ErrDuplicatedKey(s string) error {
	return fmt.Errorf("Duplicated Key on: %s", s)
}

func ErrInvalidRequest(err error, s string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     s,
		ErrorText:      err.Error(),
	}
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func ErrEmptyParam(s string) error {
	return fmt.Errorf("empty_parameter: %s should not be empty", s)

}
