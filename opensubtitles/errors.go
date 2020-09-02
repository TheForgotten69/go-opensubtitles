package opensubtitles

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Response *http.Response
	Errors []string `json:"errors"`
	Status int      `json:"status"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Status, e.Errors)
}
