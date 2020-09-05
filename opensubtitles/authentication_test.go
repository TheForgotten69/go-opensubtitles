package opensubtitles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const AuthenticationServiceLoginTestdata = "../testdata/authentication/login.json"

func TestAuthenticationService_Login(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	data, err := readFileContents(AuthenticationServiceLoginTestdata)
	if err != nil {
		t.Errorf("Unable to open Authentication.Login test data file at %s", AuthenticationServiceLoginTestdata)
	}

	mux.HandleFunc("/api/v1/login", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, data)
	})
	opt := Credentials{
		"test",
		"test",
	}
	loggedIn, _, err := client.Authentication.Login(context.Background(), &opt)
	if err != nil {
		t.Errorf("Authentication.Login returned error: %v", err)
	}

	var want *LoggedIn
	err = json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Errorf("Authentication.Login test data couldn't be Unmarshal")
	}
	if !reflect.DeepEqual(loggedIn, want) {
		t.Errorf("Authentication.Login returned %+v, want %+v", loggedIn, want)
	}
}
