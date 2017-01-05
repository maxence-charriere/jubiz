package main

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	regexpSrc = regexp.MustCompile(`src="(.+?)"`)

	videoProviders = []string{
		"https://www.youtube.com",
		"https://player.vimeo.com",
		"//www.dailymotion.com/embed/video",
	}
)

type video struct {
	URL *url.URL
}

func parseVideo(tag string) (v video, err error) {
	src := regexpSrc.FindString(tag)

	if len(src) == 0 {
		err = errors.New("no src property")
		return
	}

	srcSplit := strings.Split(src, "=")
	if len(srcSplit) < 2 {
		err = fmt.Errorf("invalid src attribute: %v", src)
		return
	}

	endpoint := srcSplit[1]
	endpoint = strings.Trim(endpoint, `"`)

	if !isTrustedVideoSource(endpoint) {
		err = fmt.Errorf("not trusted video source: %v", src)
		return
	}

	v.URL, err = url.Parse(endpoint)
	return
}

func isTrustedVideoSource(endpoint string) bool {
	for _, provider := range videoProviders {
		if strings.HasPrefix(endpoint, provider) {
			return true
		}
	}
	return false
}
