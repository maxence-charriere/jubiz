package jubiz

import (
	"strings"
)

type Image struct {
	Name string
	URL  string
}

func makeImage(url string) Image {
	name := strings.TrimPrefix(url, "http://www.fubiz.net/wp-content/uploads/")
	name = strings.Replace(name, "/", "-", -1)

	return Image{
		Name: name,
		URL:  url,
	}
}
