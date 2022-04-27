package hoster

import (
	"net/http"
	"net/url"
	"regexp"
	"server/lib/crawler"
	"server/lib/schema"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type UqloadHoster struct {
	schema.MovieHoster
	*http.Client
}

func (uh *UqloadHoster) Verify(name string) bool {
	return strings.Contains("uqloadcom", name)
}

func (uh *UqloadHoster) SetData(hoster schema.MovieHoster, client *http.Client) {
	uh.MovieHoster = hoster
	uh.Client = client
}

func (uh *UqloadHoster) Video(link, referer string) *schema.MovieVideo {
	url, _ := url.Parse(link)
	response, err := uh.Do(&http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"Referer":      []string{referer},
		},
		URL: url,
	})
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return uh.video(crawler.NewElement(document.Selection))
}

func (uh *UqloadHoster) video(document *crawler.Element) *schema.MovieVideo {
	regexp := regexp.MustCompile(`https\:\/\/(\S+).mp4`)
	result := regexp.FindString(document.Content)
	return &schema.MovieVideo{Link: result, Referer: uh.URL}
}
