package errs

import "net/http"

type AppErr struct {
	Code    int `json:",omitempty"`
	Message string
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{
		Message: e.Message,
	}
}

func NewNotFoundError(msg string) *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewUnexpectedError(msg string) *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func NewValidationError(msg string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
	}
}
