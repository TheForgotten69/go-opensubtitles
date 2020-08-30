package main

import (
	"context"
	"fmt"
	"github.com/TheForgotten69/go-opensubtitles/opensubtitles"
)

func main() {
	client := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: "abdalaoe",
		Password: "abdalaoe",
	})

	client, _ = client.Connect()
	//fmt.Println(client)
	a, _, err := client.Discover.MostDownloaded(context.Background(), nil)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(a)
}
