package opensubtitles

import (
	"context"
	"net/http"
)

type DownloadService service

type DownloadOptions struct {
	//ID of the file to download
	FileID int `url:"file_id,omitempty"`
	//Format (optional, default to original. possible values: srt, sub, mpl, webvtt, dfxp, txt)
	SubFormat string `url:"sub_format,omitempty"`
	//Desired name of the returned file
	FileName string `url:"file_name,omitempty"`
	//Remove HTML tags (default false)
	StripHTML string `url:"strip_html,omitempty"`
	//Remove HTML links
	CleanupLinks string `url:"cleanup_links,omitempty"`
	//Remove ads
	RemoveAds string `url:"remove_adds,omitempty"`
	//Input FPS (advanced, default to original subtitle FPS)
	InFPS string `url:"in_fps,omitempty"`
	//Output FPS
	OutFPS string `url:"out_fps,omitempty"`
	//Timeshift (+/- time in ms or s, eg +2s or -200ms)
	Timeshift string `url:"timeshift,omitempty"`
}

type Download struct {
	Link      string `json:"link"`
	Fname     string `json:"fname"`
	Requests  int    `json:"requests"`
	Allowed   int    `json:"allowed"`
	Remaining int    `json:"remaining"`
	Message   string `json:"message"`
}

//Download file specified by an id
//OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#download-subtitle-file
func (s *DownloadService) Download(ctx context.Context, opt *DownloadOptions) (*Download, *http.Response, error) {
	u := "/api/v1/download"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("POST", u, "")
	if err != nil {
		return nil, nil, err
	}

	var download *Download
	resp, err := s.client.Do(ctx, req, &download)
	if err != nil {
		return nil, resp, err
	}
	return download, resp, nil

}