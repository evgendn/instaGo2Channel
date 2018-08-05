package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"os"
)

func downloadMedia(url string) (err error) {
	splittedURL := strings.Split(url, "/")
	filename := splittedURL[len(splittedURL) - 1]

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %s", response.Status)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}