package opensubtitles

import (
	"context"
	"net/http"
)

// DiscoverService provides access to the discover related functions
// in the OpenSubtitles API.
type DiscoverService service

// DiscoveryOptions is used for the discover API
type DiscoveryOptions struct {
	//All, or language code
	Language string `url:"language,omitempty"`
	//Type (movie or tvshow)
	Type string `url:"type,omitempty"`
}

// Discover popular subtitles, according to last 30 days downloads on opensubtitles.com.
// OpenSubtitles API docs : https://opensubtitles.stoplight.io/docs/opensubtitles-api/3a149b956fcab-most-downloaded-subtitles
func (s *DiscoverService) MostDownloaded(ctx context.Context, opt *DiscoveryOptions) (sub *Subtitles, resp *http.Response, err error) {
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

// Discover popular features on opensubtitles.com, according to last 30 days downloads.
// OpenSubtitles API docs : https://opensubtitles.stoplight.io/docs/opensubtitles-api/6d285998026d0-popular-features
func (s *DiscoverService) Popular(ctx context.Context, opt *DiscoveryOptions) (subtitles *Subtitles, resp *http.Response, err error) {
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

// Lists 60 latest uploaded subtitles
// OpenSubtitles API docs : https://api.opensubtitles.com/api/v1/discover/latest
func (s *DiscoverService) Latest(ctx context.Context, opt *DiscoveryOptions) (subtitles *Subtitles, resp *http.Response, err error) {
	u := "/api/v1/discover/latest"
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
