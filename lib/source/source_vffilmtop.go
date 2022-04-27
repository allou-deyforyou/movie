package source

import (
	"fmt"
	"net/http"
	"strings"

	"server/lib/crawler"
	"server/lib/schema"

	"github.com/PuerkitoBio/goquery"
)

type Vffilmtop struct {
	schema.MovieSource
	*http.Client
}

func (is *Vffilmtop) Verify(name string) bool {
	return strings.Contains("vffilmtop", name)
}

func (is *Vffilmtop) SetData(source schema.MovieSource, client *http.Client) {
	is.MovieSource = source
	is.Client = client
}

func (is *Vffilmtop) FilmPostList(page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.FilmListURL, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmPostList(crawler.NewElement(document.Selection))
}

func (is *Vffilmtop) filmPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.FilmPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List,
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title)
			filmList = append(filmList, schema.MoviePost{
				Image:  strings.ReplaceAll(image, "w185", "w780"),
				Source: is.URL,
				Title:  title,
				Link:   link,
			})
		})
	return filmList
}

func (is *Vffilmtop) SearchFilmPostList(query string, page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.FilmListURL, query, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmPostList(crawler.NewElement(document.Selection))
}

func (is *Vffilmtop) FilmArticle(link string) *schema.MovieFilmArticle {
	response, err := is.Get(link)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmArticle(crawler.NewElement(document.Selection))
}

func (is *Vffilmtop) filmArticle(document *crawler.Element) *schema.MovieFilmArticle {
	hosters := document.ChildTexts(is.FilmHosterListSelector[0])
	return &schema.MovieFilmArticle{Hosters: hosters}
}
