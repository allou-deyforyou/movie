package hoster

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"server/lib/crawler"
	"server/lib/schema"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type UserloadHoster struct {
	schema.MovieHoster
	*http.Client
}

func (uh *UserloadHoster) Verify(name string) bool {
	return strings.Contains("userloadco", name)
}

func (uh *UserloadHoster) SetData(hoster schema.MovieHoster, client *http.Client) {
	uh.MovieHoster = hoster
	uh.Client = client
}

func (uh *UserloadHoster) Video(link, referer string) *schema.MovieVideo {
	url, _ := url.Parse(link)
	response, err := uh.Do(&http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"Referer": []string{referer},
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

func (uh *UserloadHoster) video(document *crawler.Element) *schema.MovieVideo {
	regexp := regexp.MustCompile(`var\|\|(?P<codes>\S+)\'.split`)
	matches := regexp.FindStringSubmatch(document.Content)
	data := strings.Split(matches[regexp.SubexpIndex("codes")], "|")
	mycountry := data[6]
	morocco := data[2]

	link, _ := url.Parse(fmt.Sprintf("%s/api/request/", uh.URL))
	response, err := uh.Do(&http.Request{
		Method: http.MethodPost,
		Body: io.NopCloser(
			bytes.NewReader([]byte(url.Values{
				"mycountry": []string{mycountry},
				"morocco":   []string{morocco},
			}.Encode())),
		),
		Header: http.Header{
			"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0"},
			"Content-Type": []string{"application/x-www-form-urlencoded"},
			"Referer":      []string{uh.URL},
		},
		URL: link,
	})
	if err != nil {
		return nil
	}
	b, _ := io.ReadAll(response.Body)
	return &schema.MovieVideo{Link: string(b), Referer: uh.URL}
}
