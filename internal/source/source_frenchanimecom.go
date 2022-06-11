package source

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"yola/internal/crawler"
	"yola/internal/schema"

	"github.com/PuerkitoBio/goquery"
)

type FrenchAnimeComSource struct {
	*schema.MovieSource
	*http.Client
}

func NewFrenchAnimeComSource(source *schema.MovieSource) *FrenchAnimeComSource {
	return &FrenchAnimeComSource{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *FrenchAnimeComSource) MangaLatestPostList(page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.MangaLatestURL, page)))
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.mangaLatestPostList(crawler.NewElement(document.Selection))
}

func (is *FrenchAnimeComSource) mangaLatestPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.MangaLatestPostSelector
	mangaList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			image = is.URL + image
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
			mangaList = append(mangaList, schema.MoviePost{
				Category: schema.MovieManga,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return mangaList
}

func (is *FrenchAnimeComSource) MangaSearchPostList(query string, page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.MangaSearchURL, page)))
	log.Println(response.Status)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.mangaSearchPostList(crawler.NewElement(document.Selection))
}

func (is *FrenchAnimeComSource) mangaSearchPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.MangaSearchPostSelector
	mangaList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			image = is.URL + image
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
			mangaList = append(mangaList, schema.MoviePost{
				Category: schema.MovieManga,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return mangaList
}

func (is *FrenchAnimeComSource) MangaArticle(link string) *schema.MovieArticle {
	response, err := is.Get(link)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.mangaArticle(crawler.NewElement(document.Selection))
}

func (is *FrenchAnimeComSource) mangaArticle(document *crawler.Element) *schema.MovieArticle {
	articleSelector := is.MangaArticleSelector
	// imdb := document.ChildText(articleSelector.Imdb[0])

	var description string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.Text()), "synopsis") {
				description = strings.TrimSpace(e.ChildText(articleSelector.Date[1]))
				return false
			}
			return true
		})

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.Text()), "sortie") {
				date = strings.TrimSpace(e.ChildText(articleSelector.Date[1]))
				return false
			}
			return true
		})

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.Text()), "genre") {
				genders = append(genders, e.ChildTexts(articleSelector.Genders[0]+" a")...)
				return false
			}
			return true
		})
	videos := make([]schema.MovieVideo, 0)
	data := document.ChildText(articleSelector.Hosters[0])
	for index, value := range strings.Split(data, "!") {
		if index == 0 {
			continue
		}
		videos = append(videos, schema.MovieVideo{
			SubtitleHosters: strings.Split(value, ","),
			Name:            strconv.Itoa(index),
			Hosters:         []string{},
		})
	}
	if len(genders) == 0 {
		genders = append(genders, "N/A")
	}
	return &schema.MovieArticle{
		Description: description,
		Genders:     genders,
		Videos:      videos,
		Imdb:        "N/A",
		Date:        date,
	}
}
