package opensubtitles

import (
	"context"
	"net/http"
)

// DownloadService provides access to the download related functions
// in the OpenSubtitles API.
//
// OpenSubtitles API docs: https://www.opensubtitles.com/docs/api/html/index.htm#download
type DownloadService service

// DownloadOptions contains the parameters for the DownloadService.Download
type DownloadOptions struct {
	//ID of the file to download
	FileID int `url:"file_id,omitempty"`
	//Format (optional, default to original. possible values: srt, sub, mpl, webvtt, dfxp, txt)
	SubFormat string `url:"sub_format,omitempty"`
	//Desired name of the returned file
	FileName string `url:"file_name,omitempty"`
	//Input FPS (advanced, default to original subtitle FPS)
	InFPS string `url:"in_fps,omitempty"`
	//Output FPS
	OutFPS string `url:"out_fps,omitempty"`
	//Timeshift (+/- time in ms or s, eg +2s or -200ms)
	Timeshift string `url:"timeshift,omitempty"`
	//(1/0) set subtitle file headers to "application/force-download"
	ForceDownload bool `url:"force_download,omitempty"`
}

// Download is the return of DownloadService.Download
type Download struct {
	Link         string `json:"link"`
	FileName     string `json:"file_name"`
	Requests     int    `json:"requests"`
	Remaining    int    `json:"remaining"`
	Message      string `json:"message"`
	ResetTime    string `json:"reset_time"`
	ResetTimeUTC string `json:"reset_time_utc"`
}

// Download file specified by an id
// OpenSubtitles API docs : https://www.opensubtitles.com/docs/api/html/index.htm#download-subtitle-file
func (s *DownloadService) Download(ctx context.Context, opt *DownloadOptions) (download *Download, resp *http.Response, err error) {
	u := "/api/v1/download"
	u, err = addOptions(u, opt)
	if err != nil {
		return nil, nil, nil
	}
	req, err := s.client.NewRequest("POST", u, "")
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.Do(ctx, req, &download)
	if err != nil {
		return nil, resp, err
	}
	return
}
