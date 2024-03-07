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
	FindServiceFeaturesTestData  = "../testdata/find/features.json"
	FindServiceSubtitlesTestData = "../testdata/find/subtitles.json"
)

func TestFindService_Feature(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(FindServiceFeaturesTestData)
	if err != nil {
		t.Errorf("Unable to open FindService.Feature test data file at %s", FindServiceFeaturesTestData)
	}

	mux.HandleFunc("/api/v1/features", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := FeatureOptions{
		FeatureID: 0,
		ImdbID:    "",
		TmdbID:    "",
	}
	feature, _, err := client.Find.Features(context.Background(), &opt)
	if err != nil {
		t.Errorf("FindService.Feature returned error: %v", err)
	}

	var want *Features
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("FindService.Feature test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(feature, want) {
		t.Errorf("FindService.Feature returned %+v, want %+v", feature, want)
	}
}

func TestFindService_Find(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(FindServiceSubtitlesTestData)
	if err != nil {
		t.Errorf("Unable to open FindService.Find test data file at %s", FindServiceSubtitlesTestData)
	}

	mux.HandleFunc("/api/v1/subtitles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := SubtitlesOptions{
		ID:                646193,
		ImdbID:            0,
		TmdbID:            0,
		Type:              "",
		Query:             "",
		Languages:         "en",
		MovieHash:         "",
		UserID:            "",
		HearingImpaired:   "",
		TrustedSources:    "",
		MachineTranslated: "",
		AiTranslated:      "",
		OrderBy:           "",
		OrderDirection:    "",
	}
	find, _, err := client.Find.Subtitles(context.Background(), &opt)
	if err != nil {
		t.Errorf("FindService.Find returned error: %v", err)
	}

	var want *Subtitles
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("FindService.Find test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(find, want) {
		t.Errorf("FindService.Find returned %+v,\n want %+v", find, want)
	}
}
