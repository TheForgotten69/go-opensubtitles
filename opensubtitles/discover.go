package opensubtitles

import (
	"context"
	"net/http"
)

type DiscoverService service

//List most downloaded subtitles
func (s *DiscoverService) MostDownloaded(ctx context.Context, opt *SubtitlesOptions) (*Subtitles, *http.Response, error) {
	u := "/api/v1/discover/most_downloaded"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var subtitles *Subtitles
	resp, err := s.client.Do(ctx, req, &subtitles)
	if err != nil {
		return nil, resp, err
	}
	return subtitles, resp, nil

}

//List movies with most subtitles downloads
func (s *DiscoverService) Popular(ctx context.Context, opt *SubtitlesOptions) (*Data, *http.Response, error) {
	u := "/api/v1/discover/most_downloaded"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var subtitles *Data
	resp, err := s.client.Do(ctx, req, &subtitles)
	if err != nil {
		return nil, resp, err
	}
	return subtitles, resp, nil

}