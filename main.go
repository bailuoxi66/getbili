package main

import (
	"fmt"
	"get-bili/downloader"
	myfmt "get-bili/fmt"
)

func main() {
	fmt.Println("hello world")
	myfmt.Logger.Println("hello world...")

	request := downloader.InfoRequest{Bvids: []string{"BV1Ff4y187q9", "BV1wE411d7th"}}
	response, err := downloader.BacthDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	for _, resp := range response.Infos {
		myfmt.Logger.Printf("title:%s\n", resp.Data.Title)
	}
	myfmt.Logger.Println(response)
}
