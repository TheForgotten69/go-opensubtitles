package opensubtitles

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestDownloadService_Download(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *DownloadOptions
	}
	tests := []struct {
		name         string
		s            DownloadService
		args         args
		wantDownload *Download
		wantResp     *http.Response
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDownload, gotResp, err := tt.s.Download(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDownload, tt.wantDownload) {
				t.Errorf("Download() gotDownload = %v, want %v", gotDownload, tt.wantDownload)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Download() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
