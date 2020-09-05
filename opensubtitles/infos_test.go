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
	InfoServiceFormatsTestData   = "../testdata/infos/formats.json"
	InfoServiceLanguagesTestData = "../testdata/infos/languages.json"
	InfoServiceUserTestData      = "../testdata/infos/user.json"
)

func TestInfoService_Formats(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(InfoServiceFormatsTestData)
	if err != nil {
		t.Errorf("Unable to open InfoService.Formats test data file at %s", InfoServiceFormatsTestData)
	}

	mux.HandleFunc("/api/v1/infos/formats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	formats, _, err := client.Info.Formats(context.Background())
	if err != nil {
		t.Errorf("InfoService.Formats returned error: %v", err)
	}

	var want *Formats
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("InfoService.Formats test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(formats, want) {
		t.Errorf("InfoService.Formats returned %+v, want %+v", formats, want)
	}
}

func TestInfoService_Languages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(InfoServiceLanguagesTestData)
	if err != nil {
		t.Errorf("Unable to open InfoService.Languages test data file at %s", InfoServiceLanguagesTestData)
	}

	mux.HandleFunc("/api/v1/infos/languages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	languages, _, err := client.Info.Languages(context.Background())
	if err != nil {
		t.Errorf("InfoService.Languages returned error: %v", err)
	}

	var want *Languages
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("InfoService.Languages test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(languages, want) {
		t.Errorf("InfoService.Languages returned %+v, want %+v", languages, want)
	}
}

func TestInfoService_User(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(InfoServiceUserTestData)
	if err != nil {
		t.Errorf("Unable to open InfoService.User test data file at %s", InfoServiceUserTestData)
	}

	mux.HandleFunc("/api/v1/infos/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	user, _, err := client.Info.User(context.Background())
	if err != nil {
		t.Errorf("InfoService.User returned error: %v", err)
	}

	var want *User
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("InfoService.User test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("InfoService.User returned %+v, want %+v", user, want)
	}
}
