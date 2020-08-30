package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthenticationService_Login(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/login", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
  "user": {
    "jti": "c52d6a50-087d-4fb2-8cdf-1ae504b4cc54",
    "allowed_downloads": 100,
    "level": "Sub leecher",
    "user_id": 1,
    "ext_installed": false,
    "vip": false
  },
  "token": "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJjNTJkNmE1MC0wODdkLTRmYjItOGNkZi0xYWU1MDRiNGNjNTQiLCJhbGxvd2VkX2Rvd25sb2FkcyI6MTAwLCJsZXZlbCI6IlN1YiBsZWVjaGVyIiwidXNlcl9pZCI6MSwiZXh0X2luc3RhbGxlZCI6ZmFsc2UsInZpcCI6ZmFsc2UsInN1YiI6IjEiLCJzY3AiOiJ1c2VyIiwiYXVkIjpudWxsLCJpYXQiOjE1OTg3MDExNTQsImV4cCI6MTU5ODc4NzU1NH0.zH-RtpuDFxOF90WW1NUjdasQbxfK7Mug7FroerNJbb4",
  "status": 200
}`)
	})
	opt := Credentials{
		"test",
		"test",
	}
	loggedIn, _, err := client.Authentication.Login(context.Background(), &opt)
	if err != nil {
		t.Errorf("Authentication.Login returned error: %v", err)
	}
	want := LoggedIn{
		User: UserData{
			Jti:                "c52d6a50-087d-4fb2-8cdf-1ae504b4cc54",
			AllowedDownloads:   100,
			Level:              "Sub leecher",
			UserID:             1,
			ExtInstalled:       false,
			Vip:                false,
			RemainingDownloads: 0,
		},
		Token:  "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJjNTJkNmE1MC0wODdkLTRmYjItOGNkZi0xYWU1MDRiNGNjNTQiLCJhbGxvd2VkX2Rvd25sb2FkcyI6MTAwLCJsZXZlbCI6IlN1YiBsZWVjaGVyIiwidXNlcl9pZCI6MSwiZXh0X2luc3RhbGxlZCI6ZmFsc2UsInZpcCI6ZmFsc2UsInN1YiI6IjEiLCJzY3AiOiJ1c2VyIiwiYXVkIjpudWxsLCJpYXQiOjE1OTg3MDExNTQsImV4cCI6MTU5ODc4NzU1NH0.zH-RtpuDFxOF90WW1NUjdasQbxfK7Mug7FroerNJbb4",
		Status: 200,
	}
	if !reflect.DeepEqual(loggedIn, &want) {
		t.Errorf("Authentication.Login returned %+v, want %+v", loggedIn, want)
	}
}
