package source

import (
	"fmt"
	"net/http"
	"strings"

	"server/lib/crawler"
	"server/lib/schema"

	"github.com/PuerkitoBio/goquery"
)

type IllimitestreamingcoSource struct {
	schema.MovieFilmSource
	*http.Client
}

func (is *IllimitestreamingcoSource) Verify(name string) bool {
	return strings.Contains("illimitestreamingco", name)
}

func (is *IllimitestreamingcoSource) SetData(source schema.MovieFilmSource, client *http.Client) {
	is.MovieFilmSource = source
	is.Client = client
}

func (is *IllimitestreamingcoSource) FilmPostList(page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.LatestFilmsURL, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmPostList(crawler.NewElement(document.Selection))
}

func (is *IllimitestreamingcoSource) filmPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.FilmPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			filmList = append(filmList, schema.MoviePost{
				Image:  strings.ReplaceAll(image, "w185", "w780"),
				Source: is.Name,
				Title:  title,
				Link:   link,
			})
		})
	return filmList
}

func (is *IllimitestreamingcoSource) SearchFilmPostList(query string, page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.SearchFilmsURL, page, query)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmPostList(crawler.NewElement(document.Selection))
}

func (is *IllimitestreamingcoSource) FilmArticle(link string) *schema.MovieFilmArticle {
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

func (is *IllimitestreamingcoSource) filmArticle(document *crawler.Element) *schema.MovieFilmArticle {
	articleSelector := is.MovieFilmArticleSelector
	hosters := document.ChildTexts(articleSelector.Hosters[0])
	description := document.ChildText(articleSelector.Description[0])

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(e.ChildText("strong"), "Genre") {
				genders = append(genders, e.ChildTexts("a")...)
				return false
			}
			return true
		})

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *crawler.Element) bool {
			fmt.Println(e.Content)

			if strings.Contains(e.ChildText("strong"), "Année") {
				date = e.ChildText("a")
				return false
			}
			return true
		})
	imdb := document.ChildText(articleSelector.Imdb[0])
	return &schema.MovieFilmArticle{
		Description: description,
		Hosters:     hosters,
		Genders:     genders,
		Date:        date,
		Imdb:        imdb,
	}
}
