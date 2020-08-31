package opensubtitles

import "time"

//Subtitles
type Subtitles struct {
	TotalPages int    `json:"total_pages"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Data       []Data `json:"data"`
}
//Uploader
type Uploader struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}
//FeatureDetails
type FeatureDetails struct {
	FeatureID   int    `json:"feature_id"`
	FeatureType string `json:"feature_type"`
	Year        int    `json:"year"`
	Title       string `json:"title"`
	MovieName   string `json:"movie_name"`
	ImdbID      int    `json:"imdb_id"`
	TmdbID      int    `json:"tmdb_id"`
}
//RelatedLinks
type RelatedLinks struct {
	Label  string `json:"label"`
	URL    string `json:"url"`
	ImgURL string `json:"img_url"`
}
//Files
type Files struct {
	ID       int    `json:"id"`
	CdNumber int    `json:"cd_number"`
	FileName string `json:"file_name"`
}
//Attributes
type Attributes struct {
	Language          string         `json:"language"`
	DownloadCount     int            `json:"download_count"`
	NewDownloadCount  int            `json:"new_download_count"`
	HearingImpaired   bool           `json:"hearing_impaired"`
	Hd                bool           `json:"hd"`
	Format            interface{}    `json:"format"`
	Fps               float64        `json:"fps"`
	Votes             int            `json:"votes"`
	Points            int            `json:"points"`
	Ratings           float64        `json:"ratings"`
	FromTrusted       bool           `json:"from_trusted"`
	AutoTranslation   bool           `json:"auto_translation"`
	AiTranslated      bool           `json:"ai_translated"`
	MachineTranslated interface{}    `json:"machine_translated"`
	UploadDate        time.Time      `json:"upload_date"`
	FileHashes        []string       `json:"file_hashes"`
	Release           string         `json:"release"`
	Comments          string         `json:"comments"`
	LegacySubtitleID  int            `json:"legacy_subtitle_id"`
	Uploader          Uploader       `json:"uploader"`
	FeatureDetails    FeatureDetails `json:"feature_details"`
	URL               string         `json:"url"`
	RelatedLinks      RelatedLinks   `json:"related_links"`
	Files             []Files        `json:"files"`
	SubtitleID        string         `json:"subtitle_id"`
}
//Data
type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}
//SubtitlesOptions is used for the discover API
type SubtitlesOptions struct {
	//All, or language code
	Language string `url:"language,omitempty"`
	//Type (movie or tvshow)
	Type string `url:"type,omitempty"`
}

//Shows represent a list of tv show or movies
type Shows struct {
	Data []Show `json:"data"`
}
//SubtitlesCounts return the subtitle count by language
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
//ShowAttributes represent the default show attributes
type ShowAttributes struct {
	Title           string          `json:"title"`
	OriginalTitle   string          `json:"original_title"`
	ImdbID          int             `json:"imdb_id"`
	TmdbID          int             `json:"tmdb_id"`
	FeatureID       string          `json:"feature_id"`
	Year            string          `json:"year"`
	TitleAka        []string        `json:"title_aka"`
	SubtitlesCounts SubtitlesCounts `json:"subtitles_counts"`
	URL             string          `json:"url"`
	ImgURL          string          `json:"img_url"`
}
//Show represent a movie or a tv show
type Show struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes ShowAttributes `json:"attributes"`
}
