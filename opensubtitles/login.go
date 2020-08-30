package opensubtitles

import (
	"context"
	"net/http"
)

// LoginService provides access to the login related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#authentication
type LoginService service

type LoggedIn struct {
	User   User   `json:"user"`
	Token  string `json:"token"`
	Status int    `json:"status"`
}


func (s *LoginService) Login(ctx context.Context, opt *Credentials) (*LoggedIn, *http.Response, error) {
	u := "/api/v1/login"
	payload := "{\n  \"username\": \"abdalaoe\",\n  \"password\": \"abdalaoe\"\n}"//fmt.Sprintf(`{"username": "%s","password": "%s"}`, opt.Username, opt.Password)
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