package opensubtitles

type InfoService service

type User struct {
	Jti              string `json:"jti"`
	AllowedDownloads int    `json:"allowed_downloads"`
	Level            string `json:"level"`
	UserID           int    `json:"user_id"`
	ExtInstalled     bool   `json:"ext_installed"`
	Vip              bool   `json:"vip"`
}
