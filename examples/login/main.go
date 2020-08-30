package main

import (
	"github.com/TheForgotten69/go-opensubtitles/opensubtitles"
	"log"
)

func main() {
	client := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: "test",
		Password: "test",
	})

	client, err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
