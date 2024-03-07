package opensubtitles

import "time"

// Subtitles is all the subtitles for a given file/show
type Subtitles struct {
	TotalPages int            `json:"total_pages"`
	TotalCount int            `json:"total_count"`
	Page       int            `json:"page"`
	Items      []SubtitleData `json:"data"`
}

// SubtitleData is not documented currently
type SubtitleData struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes SubtitleAttributes `json:"attributes"`
}

// SubtitleAttributes for a given Subtitle
type SubtitleAttributes struct {
	SubtitleID        string         `json:"subtitle_id"`
	Language          string         `json:"language"`
	DownloadCount     int            `json:"download_count"`
	NewDownloadCount  int            `json:"new_download_count"`
	HearingImpaired   bool           `json:"hearing_impaired"`
	HD                bool           `json:"hd"`
	FPS               float64        `json:"fps"`
	Votes             int            `json:"votes"`
	Points            int            `json:"points"`
	Ratings           float64        `json:"ratings"`
	FromTrusted       bool           `json:"from_trusted"`
	ForeignPartsOnly  bool           `json:"foreign_parts_only"`
	AiTranslated      bool           `json:"ai_translated"`
	MachineTranslated bool           `json:"machine_translated"`
	UploadDate        time.Time      `json:"upload_date"`
	Release           string         `json:"release"`
	Comments          string         `json:"comments"`
	LegacySubtitleID  int            `json:"legacy_subtitle_id"`
	Uploader          Uploader       `json:"uploader"`
	FeatureDetails    FeatureDetails `json:"feature_details"`
	URL               string         `json:"url"`
	RelatedLinks      []RelatedLinks `json:"related_links"`
	Files             []Files        `json:"files"`
}

// Uploader returns basic information about the uploader
type Uploader struct {
	UploaderID int    `json:"uploader_id"`
	Name       string `json:"name"`
	Rank       string `json:"rank"`
}

// FeatureDetails provides the IMDB and TMDB ID among other basic description details
type FeatureDetails struct {
	FeatureID   int    `json:"feature_id"`
	FeatureType string `json:"feature_type"`
	Year        int    `json:"year"`
	Title       string `json:"title"`
	MovieName   string `json:"movie_name"`
	ImdbID      int    `json:"imdb_id"`
	TmdbID      int    `json:"tmdb_id"`
}

// RelatedLinks is not documented currently
type RelatedLinks struct {
	Label  string `json:"label"`
	URL    string `json:"url"`
	ImgURL string `json:"img_url"`
}

// Files is not documented currently
type Files struct {
	FileID   int    `json:"file_id"`
	CdNumber int    `json:"cd_number"`
	FileName string `json:"file_name"`
}
