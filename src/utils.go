package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"os"
	"hash/crc32"
)

func downloadMedia(url string) (err error) {
	file, err := os.Create(getNameFromURL(url))
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

func getNameFromURL(url string) string {
	splittedURL := strings.Split(url, "/")
	filename := splittedURL[len(splittedURL) - 1]
	return filename
}

func hashingName(url string) uint32 {
	name := getNameFromURL(url)
	hashedName := crc32.ChecksumIEEE([]byte(name))
	return hashedName
}

func getFileFormat(filename string) string {
	splittedName := strings.Split(filename, ".")
	fileFormat := splittedName[len(splittedName) - 1]
	return fileFormat
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
