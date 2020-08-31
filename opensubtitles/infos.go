package opensubtitles

import (
	"context"
	"net/http"
)

// InfoService provides access to the info related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#infos
type InfoService service

//UserData represent all the information of a User
type UserData struct {
	Jti                string `json:"jti"`
	AllowedDownloads   int    `json:"allowed_downloads"`
	Level              string `json:"level"`
	UserID             int    `json:"user_id"`
	ExtInstalled       bool   `json:"ext_installed"`
	Vip                bool   `json:"vip"`
	RemainingDownloads int    `json:"remaining_downloads"`
}

//Languages contains all the LanguagesData
type Languages struct {
	Data []LanguagesData `json:"data"`
}

//LanguagesData associates a code to a name
type LanguagesData struct {
	LanguageCode string `json:"language_code"`
	LanguageName string `json:"language_name"`
}

//Languages get the languages used on opensubtitles and their codes
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#get-the-languages-table
func (s *InfoService) Languages(ctx context.Context) (*Languages, *http.Response, error) {
	u := "/api/v1/infos/languages"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var languages *Languages
	resp, err := s.client.Do(ctx, req, &languages)
	if err != nil {
		return nil, resp, err
	}
	return languages, resp, nil
}

type Formats struct {
	Data FormatsData `json:"data"`
}
type FormatsData struct {
	OutputFormats []string `json:"output_formats"`
}

//Formats return a list of subtitles formats that can be processed by our system
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#list-subtitle-formats
func (s *InfoService) Formats(ctx context.Context) (*Formats, *http.Response, error) {
	u := "/api/v1/infos/formats"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var formats *Formats
	resp, err := s.client.Do(ctx, req, &formats)
	if err != nil {
		return nil, resp, err
	}
	return formats, resp, nil
}

type User struct {
	Data Data `json:"data"`
}

//User get user ID, level, total and remaining download quota
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#get-user-data
func (s *InfoService) User(ctx context.Context) (*User, *http.Response, error) {
	u := "/api/v1/infos/user"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var user *User
	resp, err := s.client.Do(ctx, req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
