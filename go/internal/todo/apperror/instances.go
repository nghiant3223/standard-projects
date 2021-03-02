package apperror

import "net/http"

var (
	ErrNotFound = Error{
		Text:       "todo not found",
		StatusCode: http.StatusNotFound,
	}
	ErrInvalid = Error{
		Text:       "invalid todo",
		StatusCode: http.StatusBadRequest,
	}
)
