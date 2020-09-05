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
	FindServiceFeatureTestData = "../testdata/find/feature.json"
	FindServiceFindTestData    = "../testdata/find/find.json"
)

func TestFindService_Feature(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(FindServiceFeatureTestData)
	if err != nil {
		t.Errorf("Unable to open FindService.Feature test data file at %s", FindServiceFeatureTestData)
	}

	mux.HandleFunc("/api/v1/find/feature", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := FeatureOptions{
		ID:     "",
		ImdbID: "",
		TmdbID: "",
	}
	feature, _, err := client.Find.Feature(context.Background(), &opt)
	if err != nil {
		t.Errorf("FindService.Feature returned error: %v", err)
	}

	var want *Feature
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

	data, err := readFileContents(FindServiceFeatureTestData)
	if err != nil {
		t.Errorf("Unable to open FindService.Find test data file at %s", FindServiceFeatureTestData)
	}

	mux.HandleFunc("/api/v1/find", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := FindOptions{
		ID:                "646193",
		ImdbID:            "",
		TmdbID:            "",
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
	find, _, err := client.Find.Find(context.Background(), &opt)
	if err != nil {
		t.Errorf("FindService.Find returned error: %v", err)
	}

	var want *Feature
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("FindService.Find test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(find, want) {
		t.Errorf("FindService.Find returned %+v, want %+v", find, want)
	}
}
