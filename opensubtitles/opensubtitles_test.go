package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

var ctx = context.Background()

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/api/v1/access_token", func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"access_token": "token1",
			"token_type": "bearer",
			"expires_in": 3600,
			"scope": "*"
		}`
		w.Header().Add(headerContentType, mediaTypeJSON)
		fmt.Fprint(w, response)
	})

	client, _ := NewClient(nil, "", Credentials{
		Username: "test",
		Password: "test",
	},
	)

	return client, mux, server.Close
}