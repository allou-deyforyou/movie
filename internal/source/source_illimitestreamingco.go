package source

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"yola/internal/crawler"
	"yola/internal/schema"

	"github.com/PuerkitoBio/goquery"
)

type Illimitestreamingco struct {
	*schema.MovieSource
	*http.Client
}

func NewIllimitestreamingco(source *schema.MovieSource) *Illimitestreamingco {
	return &Illimitestreamingco{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *Illimitestreamingco) FilmLatestPostList(page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.FilmLatestURL, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmLatestPostList(crawler.NewElement(document.Selection))
}

func (is *Illimitestreamingco) filmLatestPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.FilmLatestPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])

			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
			filmList = append(filmList, schema.MoviePost{
				Category: schema.MovieFilm,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return filmList
}

func (is *Illimitestreamingco) FilmSearchPostList(query string, page int) []schema.MoviePost {
	response, err := is.Get(
		fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.FilmSearchURL, page, query)),
	)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmSearchPostList(crawler.NewElement(document.Selection))
}

func (is *Illimitestreamingco) filmSearchPostList(document *crawler.Element) []schema.MoviePost {
	if !strings.Contains(strings.ToLower(document.ChildText(".movies-list-wrap .ml-title")), "recherche") {
		return nil
	}
	selector := is.FilmSearchPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
			if strings.Contains(strings.ToLower(element.ChildText(".jtip-bottom")), "film") {
				filmList = append(filmList, schema.MoviePost{
					Category: schema.MovieFilm,
					Source: is.Name,
					Image:  image,
					Title:  title,
					Link:   link,
				})
			}
		})
	return filmList
}

func (is *Illimitestreamingco) FilmArticle(link string) *schema.MovieArticle {
	response, err := is.Get(link)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmArticle(crawler.NewElement(document.Selection))
}

func (is *Illimitestreamingco) filmArticle(document *crawler.Element) *schema.MovieArticle {
	articleSelector := is.FilmArticleSelector
	description := document.ChildText(articleSelector.Description[0])
	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("strong")), "genre") {
				genders = append(genders, e.ChildTexts("a")...)
				return false
			}
			return true
		})

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("strong")), "année") {
				date = strings.TrimSpace(e.ChildText("a"))
				return false
			}
			return true
		})

	var imdb string
	document.ForEachWithBreak(articleSelector.Imdb[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("strong")), "imdb") {
				imdb = strings.TrimSpace(e.ChildText("span"))
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)
	for _, videoSelector := range articleSelector.Videos {
		subtitleHosters := make([]string, 0)
		hosters := make([]string, 0)
		document.ForEach(videoSelector.Hosters[0],
			func(i int, e *crawler.Element) {
				id := e.ChildAttribute("a", "href")
				if id != "" {
					link := document.ChildText(id)
					if strings.TrimSpace(strings.ToLower(e.ChildText("span:nth-child(2) h6"))) == "vf" {
						hosters = append(hosters, link)
					} else {
						subtitleHosters = append(subtitleHosters, link)
					}
				}
			})
		videos = append(videos, schema.MovieVideo{
			SubtitleHosters: subtitleHosters,
			Hosters:         hosters,
			Name:            "Film",
		})
	}
	if len(genders) == 0 {
		genders = append(genders, "N/A")
	}
	return &schema.MovieArticle{
		Description: description,
		Genders:     genders,
		Videos:      videos,
		Imdb:        imdb,
		Date:        date,
	}
}
