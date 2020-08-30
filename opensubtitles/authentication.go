package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
)

// LoginService provides access to the login related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#authentication
type AuthenticationService service

type LoggedIn struct {
	User   UserData `json:"user"`
	Token  string   `json:"token"`
	Status int      `json:"status"`
}

//This endpoint provides an authentication token to the rest of the API
//
//The response will return an token which should be included in all API requests to the server in a header that looks like the following:
//
//Authorization: your-auth-token
func (s *AuthenticationService) Login(ctx context.Context, opt *Credentials) (*LoggedIn, *http.Response, error) {
	u := "/api/v1/login"
	payload := fmt.Sprintf("{\n  \"username\": \"%s\",\n  \"password\": \"%s\"\n}", opt.Username, opt.Password)
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var loggedIn *LoggedIn
	resp, err := s.client.Do(ctx, req, &loggedIn)
	if err != nil {
		return nil, resp, err
	}
	return loggedIn, resp, nil

}
