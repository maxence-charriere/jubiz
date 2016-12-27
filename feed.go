package main

import (
	"encoding/xml"
	"errors"
	"net/http"
)

const (
	feedURL = "http://feeds.feedburner.com/fubiz"
)

type feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
}

type channel struct {
	LastBuildDate string `xml:"lastBuildDate"`
	Items         []item `xml:"item"`
}

type item struct {
	Title      string   `xml:"title"`
	Link       string   `xml:"link"`
	PubDate    string   `xml:"pubDate"`
	Creator    string   `xml:"creator"`
	Categories []string `xml:"category"`
	Content    string   `xml:"encoded"`
}

func getFeed() (f feed, err error) {
	res, err := http.Get(feedURL)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(res.Status)
		return
	}

	dec := xml.NewDecoder(res.Body)
	err = dec.Decode(&f)
	return
}
