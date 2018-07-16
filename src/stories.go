package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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

	htmlBody, _ := ioutil.ReadAll(response.Body)
	htmlString := string(htmlBody)
	return htmlString
}

func parseFromHTML(htmlPage string) []string {
	var linksFromPage []string

	tokenizer := html.NewTokenizer(strings.NewReader(htmlPage))
	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.ErrorToken:
			var storiesLinks  = findOriginalStories(&linksFromPage)
			return storiesLinks 
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()
			if token.DataAtom == atom.A {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						clearURLAndInsert(&attr.Val, &linksFromPage)
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

// Links inludes url photos and videos, also video has preview.
// This function returns only photo and video, exclude preview.
func findOriginalStories(urls *[]string) []string {
	var result []string
	if len(*urls) == 1 {
		result = append(result, (*urls)[0])
	}

	for i := 0; i < len(*urls) - 1; i++ {
		index := i
		if containsVideo(&(*urls)[i], &(*urls)[i + 1]) {
			index = i + 1;			
			
			// The site give maximum 3 video links.
			offset := 1;
			for offset < 3 {
				if strings.HasSuffix((*urls)[index + offset], "mp4") {
					offset++
				} else {
					break
				}
			}
			i = index + (offset - 1)
		}
		result = append(result, (*urls)[index])
	}
	return result
}

func containsVideo(url1, url2 *string) bool {
	if strings.HasSuffix(*url1, "jpg") && strings.HasSuffix(*url2, "mp4") {
		return true
	}
	return false
}
