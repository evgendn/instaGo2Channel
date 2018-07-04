package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getStories(username string) []string {
	htmlPage :=	fetchStoriesPage(username)
	storiesLinks := parseFromHTML(htmlPage)
	return storiesLinks
}

func fetchStoriesPage(username string) string {
	baseURL := "http://zasasa.com/en/download_instagram_stories.php"
	instURL := "http://instagram.com/"

	var formData strings.Builder
	formData.WriteString(instURL)
	formData.WriteString(username)

	values := url.Values{}
	values.Set("url", formData.String())

	response, err := http.PostForm(baseURL, values)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	return bodyString
}

func parseFromHTML(html string) []string {
	var storiesLinks []string

	return storiesLinks
}
