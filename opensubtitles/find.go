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
type SubtitlesCounts struct {
	En   int `json:"en"`
	Ar   int `json:"ar"`
	PtBR int `json:"pt-BR"`
	Ro   int `json:"ro"`
	Hr   int `json:"hr"`
	Pl   int `json:"pl"`
	Nl   int `json:"nl"`
	El   int `json:"el"`
	PtPT int `json:"pt-PT"`
	Es   int `json:"es"`
	Bg   int `json:"bg"`
	ID   int `json:"id"`
	Tr   int `json:"tr"`
	He   int `json:"he"`
	Sr   int `json:"sr"`
	ZhCN int `json:"zh-CN"`
	Sv   int `json:"sv"`
	Fa   int `json:"fa"`
	Hu   int `json:"hu"`
	It   int `json:"it"`
	Cs   int `json:"cs"`
	De   int `json:"de"`
	Fi   int `json:"fi"`
	Ms   int `json:"ms"`
	Vi   int `json:"vi"`
	Ja   int `json:"ja"`
	Ko   int `json:"ko"`
	Mk   int `json:"mk"`
	Sl   int `json:"sl"`
	ZhTW int `json:"zh-TW"`
	Bn   int `json:"bn"`
	Da   int `json:"da"`
	Fr   int `json:"fr"`
	No   int `json:"no"`
	Ru   int `json:"ru"`
	Si   int `json:"si"`
	Th   int `json:"th"`
	Pm   int `json:"pm"`
	Sq   int `json:"sq"`
	Et   int `json:"et"`
	Hi   int `json:"hi"`
	Ml   int `json:"ml"`
	Sk   int `json:"sk"`
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
func (s *FindService) Feature(ctx context.Context, opt *FeatureOptions) (*Feature, *http.Response, error) {
	u := "/api/v1/find/feature"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var feature *Feature
	resp, err := s.client.Do(ctx, req, &feature)
	if err != nil {
		return nil, resp, err
	}
	return feature, resp, nil

}

//Find subtitles for a movie specified by an ID or by sending his file name and moviehash
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-subtitles-by-id
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#find-subtitles-for-a-video-file
//TODO:
func (s *FindService) Find(ctx context.Context, opt *FeatureOptions) (*Feature, *http.Response, error) {
	u := "/api/v1/find"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	var feature *Feature
	resp, err := s.client.Do(ctx, req, &feature)
	if err != nil {
		return nil, resp, err
	}
	return feature, resp, nil

}
