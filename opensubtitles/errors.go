package opensubtitles

import (
	"fmt"
	"net/http"
)

// ErrorResponse contains the http response, the list of Errors and the status
type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Message  string         `json:"message"`
	Status   int            `json:"status"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Status, e.Message)
}
