package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/TheForgotten69/go-opensubtitles/opensubtitles"
)

var username = "user"
var password = "pass"
var apiKey = "key"

func main() {
	opensubtitles.UserAgent = "SubtitlesTest v0.0.1"
	client := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: username,
		Password: password,
	}, apiKey)

	client, err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got user info from login:\n%#v", client.User)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		log.Println("Main Defer: canceling context")
		cancel()
	}()

	user, _, err := client.Info.User(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got user info from /infos/user:\n%#v", user)

	movies, _, err := client.Find.Features(ctx, &opensubtitles.FeatureOptions{Query: "the flash 2023", Type: "movie"})
	if err != nil {
		log.Fatal(err)
	}

	var movie opensubtitles.FeatureData
	if len(movies.Items) > 0 {
		log.Printf("Found movies:\n%#v", movies)
		movie = movies.Items[0]
	} else {
		log.Fatal("No movies found")
	}

	movieID, err := strconv.Atoi(movie.ID)
	if err != nil {
		log.Fatal(err)
	}
	subtitles, _, err := client.Find.Subtitles(ctx, &opensubtitles.SubtitlesOptions{ID: movieID, Languages: "ru"})
	if err != nil {
		log.Fatal(err)
	}
	for _, subtitle := range subtitles.Items {
		log.Printf("Found subtitle:\n%#v", subtitle)
	}
	if len(subtitles.Items) > 0 {
		subtitle := subtitles.Items[0]
		if len(subtitle.Attributes.Files) > 0 {
			subtitleFile := subtitle.Attributes.Files[0]
			subtitleDownload, _, err := client.Download.Download(ctx, &opensubtitles.DownloadOptions{FileID: subtitleFile.FileID})
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Download 1st subtitle:\n%#v", subtitleDownload)
		}
	}

	logoutMessage, _, err := client.Authentication.Logout(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got message from /logout:\n%#v", logoutMessage)
}
