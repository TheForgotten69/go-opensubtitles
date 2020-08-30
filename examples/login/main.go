package main

import (
	"context"
	"fmt"
	"github.com/TheForgotten69/go-opensubtitles/opensubtitles"
)

func main() {
	client, err := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: "abdalaoe",
		Password: "abdalaoe",
	})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(client)
	a, _, err := client.Discover.MostDownloaded(context.Background(), nil)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(a)
}
