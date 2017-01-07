package jubiz

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	regexpSrc      = regexp.MustCompile(`src="(.+?)"`)
	videoProviders = []string{
		"https://www.youtube.com",
		"https://player.vimeo.com",
		"//www.dailymotion.com/embed/Video",
	}
)

type Video struct {
	URL *url.URL
}

func parseVideo(tag string) (v Video, err error) {
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
		err = fmt.Errorf("not trusted Video source: %v", src)
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
