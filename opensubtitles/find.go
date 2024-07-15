package opensubtitles

import (
	"context"
	"net/http"
)

// FindService provides access to the find related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#find
type FindService service

// FeatureOptions provide the parameters for FindService.Feature
type FeatureOptions struct {
	//opensubtitles feature_id
	FeatureID int `url:"feature_id,omitempty"`
	//IMDB ID, delete leading zeroes
	ImdbID string `url:"imdb_id,omitempty"`
	//TheMovieDB ID - combine with type to avoid errors
	TmdbID string `url:"tmdb_id,omitempty"`
	//query to search, release/file name accepted
	Query string `url:"query,omitempty"`
	//empty to list all or movie, tvshow or episode.
	Type string `url:"type,omitempty"`
	//Filter by year. Can only be used in combination with a query
	Year int `url:"year,omitempty"`
}

// Features find details for a movie or tv show specified by an ID.
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-details-for-a-feature
func (s *FindService) Features(ctx context.Context, opt *FeatureOptions) (feature *Features, resp *http.Response, err error) {
	u := "/api/v1/features"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}
	resp, err = s.client.Do(ctx, req, &feature)
	if err != nil {
		return nil, resp, err
	}
	return
}

// SubtitlesOptions provide the parameters for FindService.Subtitle
type SubtitlesOptions struct {
	//Feature ID of the feature
	ID int `url:"id,omitempty"`
	//IMDB ID of the feature
	ImdbID string `url:"imdb_id,omitempty"`
	//TMDB ID of the feature
	TmdbID int `url:"tmdb_id,omitempty"`
	//movie, episode or all, (default: all)
	Type string `url:"type,omitempty"`
	//file name or text search
	Query string `url:"query,omitempty"`
	//Language code(s), coma separated (en,fr)
	Languages string `url:"languages,omitempty"`
	//Moviehash of the movie
	MovieHash string `url:"moviehash,omitempty"`
	//include, only (default: include)
	MovieHashMatch string `url:"moviehash_match,omitempty"`
	//To be used alone - for user uploads listing
	UserID string `url:"user_id,omitempty"`
	//include, exclude, only.(default: include)
	HearingImpaired string `url:"hearing_impaired,omitempty"`
	//include, only (default: include)
	TrustedSources string `url:"trusted_sources,omitempty"`
	//exclude, include (default: exclude)
	MachineTranslated string `url:"machine_translated,omitempty"`
	//exclude, include (default: exclude)
	AiTranslated string `url:"ai_translated,omitempty"`
	//Order of the returned results, accept any of above fields
	OrderBy string `url:"order_by,omitempty"`
	//Order direction of the returned results (asc, desc)
	OrderDirection string `url:"order_direction,omitempty"`
	//For Tvshows
	EpisodeNumber int `url:"episode_number"`
	//exclude, include, only (default: include)
	ForeignPartsOnly string `url:"foreign_parts_only"`
	//Results page to display
	Page int `url:"page,omitempty"`
	// For Tvshows
	ParentFeatureID int `url:"parent_feature_id,omitempty"`
	// For Tvshows
	ParentImdbID int `url:"parent_imdb_id,omitempty"`
	// For Tvshows
	ParentTmdbID int `url:"parent_tmdb_id,omitempty"`
	// For Tvshows
	SeasonNumber int `url:"season_number,omitempty"`
	// To be used alone - for user uploads listing
	UserId int `url:"user_id,omitempty"`
	// Filter by movie/episode year
	Year int `url:"year,omitempty"`
}

// Find subtitles for a movie specified by an ID or by sending his file name and moviehash
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-subtitles-by-id
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-subtitles-for-a-video-file
func (s *FindService) Subtitles(ctx context.Context, opt *SubtitlesOptions) (subtitles *Subtitles, resp *http.Response, err error) {
	u := "/api/v1/subtitles"
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
	return
}
