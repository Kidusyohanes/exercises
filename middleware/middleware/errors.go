package middleware

import (
	"net/http"
)

type HTTPError struct {
	Message    string
	StatusCode int
}

type HandlerFuncReturningError func(http.ResponseWriter, *http.Request) *HTTPError

func HandleHTTPError(handlerFunc HandlerFuncReturningError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			http.Error(w, err.Message, err.StatusCode)
			return
		}
	}
}
