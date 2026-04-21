package internal

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type Downloader struct {
	client *http.Client
}

func NewDownloader() Downloader {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return Downloader{client: client}
}

func (d Downloader) DownloadData() (string, error) {
	response, err := d.client.Get("https://www.source-fellows.com/s/testdata/cars.json")
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not load data. %s", response.Status)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
