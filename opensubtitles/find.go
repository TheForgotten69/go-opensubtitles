package opensubtitles

import (
	"context"
	"net/http"
)

type FindService service

type FeatureOptions struct {
	//Feature ID of the feature
	ID string `url:"id,omitempty"`
	//IMDB ID of the feature
	ImdbID string `url:"imdb_id,omitempty"`
	//TMDB ID of the feature
	TmdbID string `url:"tmdb_id,omitempty"`
}

type Feature struct {
	Data []FeatureData `json:"data"`
}

type AttributesFeature struct {
	Title           string          `json:"title"`
	OriginalTitle   string          `json:"original_title"`
	Year            string          `json:"year"`
	SubtitlesCounts SubtitlesCounts `json:"subtitles_counts"`
	SubtitlesCount  int             `json:"subtitles_count"`
	SeasonsCount    int             `json:"seasons_count"`
	ParentTitle     string          `json:"parent_title"`
	SeasonNumber    int             `json:"season_number"`
	EpisodeNumber   interface{}     `json:"episode_number"`
	ImdbID          int             `json:"imdb_id"`
	TmdbID          int             `json:"tmdb_id"`
	ParentImdbID    interface{}     `json:"parent_imdb_id"`
	FeatureID       string          `json:"feature_id"`
	TitleAka        []string        `json:"title_aka"`
	FeatureType     string          `json:"feature_type"`
	URL             string          `json:"url"`
	ImgURL          string          `json:"img_url"`
	Seasons         []interface{}   `json:"seasons"`
}
type FeatureData struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes AttributesFeature `json:"attributes"`
}

//Find details for a movie or tv serie specified by an ID.
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-details-for-a-feature
func (s *FindService) Feature(ctx context.Context, opt *FeatureOptions) (feature *Feature, resp *http.Response, err error) {
	u := "/api/v1/find/feature"
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

type FindOptions struct {
	//Feature ID of the feature
	ID string `url:"id,omitempty"`
	//IMDB ID of the feature
	ImdbID string `url:"imdb_id,omitempty"`
	//TMDB ID of the feature
	TmdbID string `url:"tmdb_id,omitempty"`
	//movie, episode or all, (default: all)
	Type string `url:"type,omitempty"`
	//file name or text search
	Query string `url:"query,omitempty"`
	//Language code(s), coma separated (en,fr)
	Languages string `url:"languages,omitempty"`
	//Moviehash of the movie
	MovieHash string `url:"moviehash,omitempty"`
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
}

//Find subtitles for a movie specified by an ID or by sending his file name and moviehash
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-subtitles-by-id
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-subtitles-for-a-video-file
func (s *FindService) Find(ctx context.Context, opt *FindOptions) (feature *Feature, resp *http.Response, err error) {
	u := "/api/v1/find"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}
	//TODO
	resp, err = s.client.Do(ctx, req, &feature)
	if err != nil {
		return nil, resp, err
	}
	return

}
