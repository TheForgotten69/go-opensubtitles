package main

import (
	"fmt"
	"github.com/TheForgotten69/go-opensubtitles/opensubtitles"
)

func main(){
	client, _ := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: "abdalaoe",
		Password: "abdalaoe",
	})
	fmt.Println(client)
}