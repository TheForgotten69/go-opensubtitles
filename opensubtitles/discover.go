package opensubtitles

import (
	"context"
	"net/http"
)

// DiscoverService provides access to the discover related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#discover
type DiscoverService service

//MostDownloaded list the most downloaded subtitles
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-most-downloaded-movie-subtitles
func (s *DiscoverService) MostDownloaded(ctx context.Context, opt *SubtitlesOptions) (sub *Subtitles, resp *http.Response, err error) {
	u := "/api/v1/discover/most_downloaded"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &sub)
	if err != nil {
		return nil, resp, err
	}
	return
}

//Popular list movies with most subtitles downloads
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-most-downloaded-movie-subtitles
func (s *DiscoverService) Popular(ctx context.Context, opt *SubtitlesOptions) (subtitles *Data, resp *http.Response, err error) {
	u := "/api/v1/discover/popular"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &subtitles)
	if err != nil {
		return nil, resp, err
	}
	return subtitles, resp, nil
}
