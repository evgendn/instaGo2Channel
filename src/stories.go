package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	// "fmt"
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

	htmlBody, _ := ioutil.ReadAll(response.Body)
	htmlString := string(htmlBody)
	return htmlString
}

func parseFromHTML(htmlPage string) []string {
	var storiesLinks []string

	tokenizer := html.NewTokenizer(strings.NewReader(htmlPage))
	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.ErrorToken:
			return storiesLinks
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()
			if token.DataAtom == atom.A {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						clearURLAndInsert(&attr.Val, &storiesLinks)
					}
				}
			}
		}	
	}	
}

func clearURLAndInsert(rawURL *string, slice *[]string) {
	if strings.HasPrefix(*rawURL, "https") {
		url := strings.Split(*rawURL,  "?")[0]
		*slice = append(*slice, url)
	}
}