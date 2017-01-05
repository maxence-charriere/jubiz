package main

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"time"

	"html"

	"github.com/murlokswarm/log"
)

var (
	twitterURLs = map[string]string{
		"Romain":   "https://twitter.com/romaincolin",
		"Adrien":   "https://twitter.com/alepage",
		"Donnia":   "https://twitter.com/donnnia",
		"Daniella": "https://twitter.com/fubiz",
		"Léa":      "https://twitter.com/leabenatar",
		"Hoel":     "https://twitter.com/fubiz",
		"Katia":    "https://twitter.com/fubiz",
		"Apolline": "https://twitter.com/fubiz",
	}

	regexpImg      = regexp.MustCompile(`<img\s(.*?)/>`)
	regexpIframe   = regexp.MustCompile(`<iframe(.*?)/iframe>`)
	regexpEmptyA   = regexp.MustCompile(`<a(.*?)>\s*</a>`)
	regexpEmptyP   = regexp.MustCompile(`<p>\s*</p>`)
	regexpClosingP = regexp.MustCompile(`</p>(\s*</p>\s*)+`)
	regexpImgSrc   = regexp.MustCompile(`src="(.+?)\.(.{3,4})"`)
)

type article struct {
	Title      string
	URL        string
	PubDate    time.Time
	Author     author
	Text       string
	Categories []string
	Images     []image
	Videos     []video
	Read       bool
}

type articleList []article

type author struct {
	Name       string
	TwitterURL string
}

func makeArticle(i item) article {
	pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", i.PubDate)
	if err != nil {
		log.Error(err)
	}

	content := html.UnescapeString(i.Content)
	return article{
		Title:      i.Title,
		URL:        i.Link,
		PubDate:    pubDate,
		Author:     makeAuthor(i.Creator),
		Text:       makeArticleText(content),
		Categories: i.Categories,
		Images:     makeArticleImages(i.Content),
		Videos:     makeArticleVideos(i.Content),
	}
}

func makeArticleText(content string) string {
	text := regexpImg.ReplaceAllString(content, "")
	text = regexpIframe.ReplaceAllString(text, "")
	text = regexpEmptyA.ReplaceAllString(text, "")
	text = regexpEmptyP.ReplaceAllString(text, "")
	text = regexpClosingP.ReplaceAllString(text, "</p>")
	text = strings.Replace(text, "&", "&amp;", -1)
	return strings.TrimSpace(text)
}

func makeArticleVideos(content string) (videos []video) {
	videoTags := regexpIframe.FindAllString(content, -1)

	for _, tag := range videoTags {
		v, err := parseVideo(tag)
		if err != nil {
			continue
		}

		videos = append(videos, v)
	}
	return
}

func makeArticleImages(content string) (images []image) {
	imgTags := regexpImg.FindAllString(content, -1)

	for _, tag := range imgTags {
		if !strings.Contains(tag, "http://www.fubiz.net/wp-content/uploads/") {
			continue
		}

		src := regexpImgSrc.FindString(tag)
		if len(src) == 0 {
			continue
		}

		srcSplit := strings.Split(src, "=")
		if len(srcSplit) < 2 {
			continue
		}

		url := srcSplit[1]
		url = strings.Trim(url, `"`)

		if len(url) == 0 || strings.Contains(url, "-150x150.") {
			continue
		}

		img := makeImage(url)
		images = append(images, img)
	}
	return
}

func makeArticlesFromFeed(f feed) (articles articleList) {
	for _, i := range f.Channel.Items {
		articles = append(articles, makeArticle(i))
	}
	return
}

func (list articleList) Len() int {
	return len(list)
}

func (list articleList) Less(i, j int) bool {
	return list[i].PubDate.After(list[j].PubDate)
}

func (list articleList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func mergeArticleLists(base, new articleList) articleList {
	m := map[string]article{}

	for _, a := range base {
		m[a.URL] = a
	}

	for _, a := range new {
		if art, ok := m[a.URL]; ok {
			a.Read = art.Read
		}

		m[a.URL] = a
	}

	list := make(articleList, 0, len(m))

	for _, v := range m {
		list = append(list, v)
	}
	return list
}

func readArticles(name string) (a articleList, err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&a)
	return
}

func saveArticles(name string, a articleList) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	return enc.Encode(a)
}

func makeAuthor(name string) author {
	twitterURL, ok := twitterURLs[name]
	if !ok {
		twitterURL = "https://twitter.com/fubiz"
	}

	return author{
		Name:       name,
		TwitterURL: twitterURL,
	}
}
