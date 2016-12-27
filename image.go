package main

import (
	"strings"
)

type image struct {
	Name string
	URL  string
}

func makeImage(url string) image {
	name := strings.TrimPrefix(url, "http://www.fubiz.net/wp-content/uploads/")
	name = strings.Replace(name, "/", "-", -1)

	return image{
		Name: name,
		URL:  url,
	}
}
