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
	SearchServiceMovieTestData = "../testdata/search/movie_thematrix.json"
	SearchServiceTVTestData    = "../testdata/search/title_wakingdead.json"
	SearchServiceTitleTestData = "../testdata/search/tv_wakingdead.json"
)

func TestSearchService_Movie(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(SearchServiceMovieTestData)
	if err != nil {
		t.Errorf("Unable to open SearchService.Movie test data file at %s", SearchServiceMovieTestData)
	}

	mux.HandleFunc("/api/v1/search/movie", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := SearchOptions{
		"The Matrix",
	}
	movie, _, err := client.Search.Movie(context.Background(), &opt)
	if err != nil {
		t.Errorf("SearchService.Movie returned error: %v", err)
	}

	var want *Shows
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("SearchService.Movie test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(movie, want) {
		t.Errorf("SearchService.Movie returned %+v, want %+v", movie, want)
	}
}

func TestSearchService_TV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(SearchServiceTVTestData)
	if err != nil {
		t.Errorf("Unable to open SearchService.TV test data file at %s", SearchServiceTVTestData)
	}

	mux.HandleFunc("/api/v1/search/tv", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := SearchOptions{
		"Waking dead",
	}
	tv, _, err := client.Search.TV(context.Background(), &opt)
	if err != nil {
		t.Errorf("SearchService.TV returned error: %v", err)
	}

	var want *Shows
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("SearchService.TV test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(tv, want) {
		t.Errorf("SearchService.TV returned %+v, want %+v", tv, want)
	}
}

func TestSearchService_Title(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(SearchServiceTitleTestData)
	if err != nil {
		t.Errorf("Unable to open SearchService.Title test data file at %s", SearchServiceTitleTestData)
	}

	mux.HandleFunc("/api/v1/search/title", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})
	opt := SearchOptions{
		"Waking dead",
	}
	title, _, err := client.Search.Title(context.Background(), &opt)
	if err != nil {
		t.Errorf("SearchService.Title returned error: %v", err)
	}

	var want *Shows
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("SearchService.Title test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(title, want) {
		t.Errorf("SearchService.Title returned %+v, want %+v", title, want)
	}
}
