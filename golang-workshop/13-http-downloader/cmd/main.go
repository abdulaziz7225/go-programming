package main

import (
	"fmt"
	"log"

	"golang.source-fellows.com/training/http-downloader/internal"
)

func main() {
	dl := internal.NewDownloader()
	cars, err := dl.DownloadData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", cars)
}
