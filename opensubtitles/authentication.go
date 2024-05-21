package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
)

// AuthenticationService provides access to the login related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#authentication
type AuthenticationService service

// LoggedIn is return when the AuthenticationService.Login is returned with a success
type LoggedIn struct {
	User    UserData `json:"user"`
	Token   string   `json:"token"`
	Status  int      `json:"status"`
	BaseURL string   `json:"base_url"`
}

// Login endpoint provides an authentication token to the rest of the API
//
// The response will return an token which should be included in all API requests to the server in a header that looks like the following:
//
// Authorization: your-auth-token
func (s *AuthenticationService) Login(ctx context.Context, opt *Credentials) (loggedIn *LoggedIn, resp *http.Response, err error) {
	u := "/api/v1/login"
	//payload, _ := json.Marshal(map[string]string{"username": opt.Username, "password": opt.Password})
	payload := fmt.Sprintf("{\n  \"username\": \"%s\",\n  \"password\": \"%s\"\n}", opt.Username, opt.Password)
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &loggedIn)
	if err != nil {
		return nil, resp, err
	}
	return
}

func (s *AuthenticationService) Logout(ctx context.Context) (message *ErrorResponse, resp *http.Response, err error) {
	u := "/api/v1/logout"
	req, err := s.client.NewRequest("DELETE", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &message)
	if err != nil {
		return nil, resp, err
	}
	message.Response = resp
	return
}
