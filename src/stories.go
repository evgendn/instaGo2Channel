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

func findOriginalStories(urls *[]string) []string {
	var result []string
	if len(*urls) == 1 {
		result = append(result, (*urls)[0])
	}
	for i := 0; i < len(*urls) - 1; i++ {
		// Scary fucking shit.
		index := tryFindVideoIndex(&((*urls)[i]), &((*urls)[i + 1]), &i)
		result = append(result, (*urls)[index])
	}
	return result
}

// Video has preview above as jpg format.
// First checking two adjacent items of slice and change index by 3 because 
// site that parsed has 3 videos different resolution for one stories. 
func tryFindVideoIndex(url1, url2 *string, i *int) int {
	j := *i
	if strings.HasSuffix(*url1, "jpg") && strings.HasSuffix(*url2, "mp4") {
		j++
		*i += 3
	}
	return j
} 
