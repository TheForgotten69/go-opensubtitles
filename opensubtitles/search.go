package opensubtitles

import (
	"context"
	"net/http"
)

//OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#search
type SearchService service

//SearchOptions contains the url parameter for the SearchService
type SearchOptions struct {
	//Title
	Query string `url:"query,omitempty"`
}

//Movie search a movie by name
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#search-for-a-movie
func (s *SearchService) Movie(ctx context.Context, opt *SearchOptions) (shows *Shows, resp *http.Response, err error) {
	u := "/api/v1/search/movie"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &shows)
	if err != nil {
		return nil, resp, err
	}
	return

}

//TV search a TV serie by name
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#search-for-a-tv-serie
func (s *SearchService) TV(ctx context.Context, opt *SearchOptions) (shows *Shows, resp *http.Response, err error) {
	u := "/api/v1/search/tv"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &shows)
	if err != nil {
		return nil, resp, err
	}
	return

}

//Title search for a feature by title (tv serie or movie)
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#search-for-a-feature-by-title
func (s *SearchService) Title(ctx context.Context, opt *SearchOptions) (shows *Shows, resp *http.Response, err error) {
	u := "/api/v1/search/title"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &shows)
	if err != nil {
		return nil, resp, err
	}
	return

}
