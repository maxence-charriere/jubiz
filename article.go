package jubiz

import (
	"html"
	"regexp"
	"strings"
	"time"

	"net/url"

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

type Article struct {
	ID         string
	Title      string
	URL        *url.URL
	PubDate    time.Time
	Author     Author
	Text       string
	Categories []string
	Images     []Image
	Videos     []Video
	Read       bool
}

type ArticleList []Article

type Author struct {
	Name       string
	TwitterURL string
}

func makeArticle(i item) (a Article, err error) {
	pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", i.PubDate)
	if err != nil {
		log.Error(err)
	}

	u, err := url.Parse(i.Link)
	if err != nil {
		return
	}

	content := html.UnescapeString(i.Content)
	a = Article{
		ID:         i.Link,
		Title:      i.Title,
		URL:        u,
		PubDate:    pubDate,
		Author:     makeAuthor(i.Creator),
		Text:       makeArticleText(content),
		Categories: i.Categories,
		Images:     makeArticleImages(i.Content),
		Videos:     makeArticleVideos(i.Content),
	}
	return
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

func makeArticleVideos(content string) (videos []Video) {
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

func makeArticleImages(content string) (images []Image) {
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

func MakeArticlesFromFeed(f Feed) (Articles ArticleList) {
	for _, i := range f.Channel.Items {
		a, err := makeArticle(i)
		if err != nil {
			continue
		}
		Articles = append(Articles, a)
	}
	return
}

func (list ArticleList) Len() int {
	return len(list)
}

func (list ArticleList) Less(i, j int) bool {
	return list[i].PubDate.After(list[j].PubDate)
}

func (list ArticleList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func MergeArticleLists(base, new ArticleList) ArticleList {
	m := map[string]Article{}

	for _, a := range base {
		m[a.ID] = a
	}

	for _, a := range new {
		if art, ok := m[a.ID]; ok {
			a.Read = art.Read
		}
		m[a.ID] = a
	}

	list := make(ArticleList, 0, len(m))
	for _, v := range m {
		list = append(list, v)
	}
	return list
}

func makeAuthor(name string) Author {
	twitterURL, ok := twitterURLs[name]
	if !ok {
		twitterURL = "https://twitter.com/fubiz"
	}

	return Author{
		Name:       name,
		TwitterURL: twitterURL,
	}
}
