package opensubtitles

// Features contains the details returned by FindService.Features
type Features struct {
	Items []FeatureData `json:"data"`
}

// FeatureData contains the detail for only a given ID, be it a tv or show
type FeatureData struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes FeatureAttributes `json:"attributes"`
}

// FeatureAttributes is the detail for a given movie or tv contains in FeatureData
type FeatureAttributes struct {
	Title           string          `json:"title"`
	OriginalTitle   string          `json:"original_title"`
	Year            string          `json:"year"`
	SubtitlesCounts SubtitlesCounts `json:"subtitles_counts"`
	SubtitlesCount  int             `json:"subtitles_count"`
	SeasonsCount    int             `json:"seasons_count"`
	ParentTitle     string          `json:"parent_title"`
	SeasonNumber    int             `json:"season_number"`
	EpisodeNumber   int             `json:"episode_number"`
	ImdbID          int             `json:"imdb_id"`
	TmdbID          int             `json:"tmdb_id"`
	ParentImdbID    int             `json:"parent_imdb_id"`
	FeatureID       string          `json:"feature_id"`
	TitleAka        []string        `json:"title_aka"`
	FeatureType     string          `json:"feature_type"`
	URL             string          `json:"url"`
	ImgURL          string          `json:"img_url"`
	Seasons         []Season        `json:"seasons"`
}

type Season struct {
	SeasonNumber int       `json:"season_number"`
	Episodes     []Episode `json:"episodes"`
}

type Episode struct {
	EpisodeNumber int    `json:"episode_number"`
	Title         string `json:"title"`
	FeatureID     int    `json:"feature_id"`
	FeatureImdbID int    `json:"feature_imdb_id"`
}

// SubtitlesCounts return the subtitle count by language
type SubtitlesCounts struct {
	Pl   int `json:"pl"`
	En   int `json:"en"`
	Tr   int `json:"tr"`
	Ro   int `json:"ro"`
	Cs   int `json:"cs"`
	Es   int `json:"es"`
	PtBR int `json:"pt-BR"`
	Sl   int `json:"sl"`
	PtPT int `json:"pt-PT"`
	Sr   int `json:"sr"`
	El   int `json:"el"`
	Bg   int `json:"bg"`
	He   int `json:"he"`
	Nl   int `json:"nl"`
	Fi   int `json:"fi"`
	Fr   int `json:"fr"`
	Hu   int `json:"hu"`
	Ar   int `json:"ar"`
	Ru   int `json:"ru"`
	Hr   int `json:"hr"`
	Da   int `json:"da"`
	Et   int `json:"et"`
	Sv   int `json:"sv"`
	Sq   int `json:"sq"`
	Bs   int `json:"bs"`
	De   int `json:"de"`
	It   int `json:"it"`
	Ko   int `json:"ko"`
	No   int `json:"no"`
	Fa   int `json:"fa"`
	Sk   int `json:"sk"`
	Mk   int `json:"mk"`
	ZhCN int `json:"zh-CN"`
	Ms   int `json:"ms"`
	ZhTW int `json:"zh-TW"`
	Bn   int `json:"bn"`
	ID   int `json:"id"`
	Lt   int `json:"lt"`
	Is   int `json:"is"`
	Ja   int `json:"ja"`
	Th   int `json:"th"`
	Ca   int `json:"ca"`
	Hi   int `json:"hi"`
	Ml   int `json:"ml"`
	Mn   int `json:"mn"`
	Vi   int `json:"vi"`
}
