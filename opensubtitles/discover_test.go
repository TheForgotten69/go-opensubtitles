package opensubtitles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const (
	DiscoverServiceMostDownloadedTestData = "../testdata/discover/most_downloaded.json"
	DiscoverServicePopularTestData        = "../testdata/discover/popular.json"
)

func TestDiscoverService_MostDownloaded(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(DiscoverServiceMostDownloadedTestData)
	if err != nil {
		t.Errorf("Unable to open DiscoverService.MostDownloaded test data file at %s", DiscoverServiceMostDownloadedTestData)
	}

	mux.HandleFunc("/api/v1/discover/most_downloaded", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := SubtitlesOptions{
		"",
		"",
	}
	mostDL, _, err := client.Discover.MostDownloaded(context.Background(), &opt)
	if err != nil {
		t.Errorf("DiscoverService.MostDownloaded returned error: %v", err)
	}

	var want *Subtitles
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("DiscoverService.MostDownloaded test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(mostDL, want) {
		t.Errorf("DiscoverService.MostDownloaded returned %+v, want %+v", mostDL, want)
	}
}

func TestDiscoverService_Popular(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(DiscoverServicePopularTestData)
	if err != nil {
		t.Errorf("Unable to open DiscoverService.Popular test data file at %s", DiscoverServicePopularTestData)
	}

	mux.HandleFunc("/api/v1/discover/popular", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := SubtitlesOptions{
		"",
		"",
	}
	popular, _, err := client.Discover.Popular(context.Background(), &opt)
	if err != nil {
		t.Errorf("DiscoverService.Popular returned error: %v", err)
	}

	var want *Data
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("DiscoverService.Popular test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(popular, want) {
		t.Errorf("DiscoverService.Popular returned %+v, want %+v", popular, want)
	}
}
