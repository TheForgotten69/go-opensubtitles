package opensubtitles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const DownloadServiceDownloadTestData = "../testdata/download/download.json"

func TestDownloadService_Download(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(DownloadServiceDownloadTestData)
	if err != nil {
		t.Errorf("Unable to open DownloadService.Download test data file at %s", DownloadServiceDownloadTestData)
	}

	mux.HandleFunc("/api/v1/download", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, data)
	})
	opt := DownloadOptions{
		FileID:    186609,
		SubFormat: "srt",
		FileName:  "The.Meg.2018.BDRip.XviD.AC3-EVO",
	}
	download, _, err := client.Download.Download(context.Background(), &opt)
	if err != nil {
		t.Errorf("DownloadService.Download returned error: %v", err)
	}

	var want *Download
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("DownloadService.Download test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(download, want) {
		t.Errorf("DownloadService.Download returned %+v, want %+v", download, want)
	}
}
