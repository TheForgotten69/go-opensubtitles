package opensubtitles

import "time"

type Subtitles struct {
	TotalPages int    `json:"total_pages"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Data       []Data `json:"data"`
}
type Uploader struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}
type FeatureDetails struct {
	FeatureID   int    `json:"feature_id"`
	FeatureType string `json:"feature_type"`
	Year        int    `json:"year"`
	Title       string `json:"title"`
	MovieName   string `json:"movie_name"`
	ImdbID      int    `json:"imdb_id"`
	TmdbID      int    `json:"tmdb_id"`
}
type RelatedLinks struct {
	Label  string `json:"label"`
	URL    string `json:"url"`
	ImgURL string `json:"img_url"`
}
type Files struct {
	ID       int    `json:"id"`
	CdNumber int    `json:"cd_number"`
	FileName string `json:"file_name"`
}
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
type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type SubtitlesOptions struct {
	//All, or language code
	Language string `url:"language,omitempty"`
	//Type (movie or tvshow)
	Type string `url:"type,omitempty"`
}