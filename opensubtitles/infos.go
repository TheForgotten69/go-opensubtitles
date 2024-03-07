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

// UserData represent all the information of a User
type UserData struct {
	AllowedDownloads    int    `json:"allowed_downloads"`
	AllowedTranslations int    `json:"allowed_translations"`
	Level               string `json:"level"`
	UserID              int    `json:"user_id"`
	ExtInstalled        bool   `json:"ext_installed"`
	Vip                 bool   `json:"vip"`
	RemainingDownloads  int    `json:"remaining_downloads"`
	DownloadsCount      int    `json:"downloads_count"`
	ResetTimeUTC        string `json:"reset_time_utc"`
	ResetTime           string `json:"reset_time"`
}

// Languages contains all the LanguagesData
type Languages struct {
	Data []LanguagesData `json:"data"`
}

// LanguagesData associates a code to a name
type LanguagesData struct {
	LanguageCode string `json:"language_code"`
	LanguageName string `json:"language_name"`
}

// Languages get the languages used on opensubtitles and their codes
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#get-the-languages-table
func (s *InfoService) Languages(ctx context.Context) (languages *Languages, resp *http.Response, err error) {
	u := "/api/v1/infos/languages"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &languages)
	if err != nil {
		return nil, resp, err
	}
	return languages, resp, nil
}

// Formats contains the list of all available formats
type Formats struct {
	Data FormatsData `json:"data"`
}

// FormatsData is contain in the Formats struct
type FormatsData struct {
	OutputFormats []string `json:"output_formats"`
}

// Formats return a list of subtitles formats that can be processed by our system
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#list-subtitle-formats
func (s *InfoService) Formats(ctx context.Context) (formats *Formats, resp *http.Response, err error) {
	u := "/api/v1/infos/formats"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &formats)
	if err != nil {
		return nil, resp, err
	}
	return formats, resp, nil
}

// User is return by InfoService.User
type User struct {
	Data UserData `json:"data"`
}

// User get user ID, level, total and remaining download quota
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#get-user-data
func (s *InfoService) User(ctx context.Context) (user *User, resp *http.Response, err error) {
	u := "/api/v1/infos/user"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
