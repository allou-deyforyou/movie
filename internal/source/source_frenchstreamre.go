package source

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"yola/internal/crawler"
	"yola/internal/schema"

	"github.com/PuerkitoBio/goquery"
)

type FrenchStreamReSource struct {
	*schema.MovieSource
	*http.Client
}

func NewFrenchStreamReSource(source *schema.MovieSource) *FrenchStreamReSource {
	return &FrenchStreamReSource{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *FrenchStreamReSource) FilmLatestPostList(page int) []schema.MoviePost {
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

func (is *FrenchStreamReSource) filmLatestPostList(document *crawler.Element) []schema.MoviePost {
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

func (is *FrenchStreamReSource) FilmSearchPostList(query string, page int) []schema.MoviePost {
	response, err := is.PostForm(
		fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.FilmSearchURL, page)),
		url.Values{
			"do":           []string{"search"},
			"subaction":    []string{"search"},
			"story":        []string{query},
			"search_start": []string{strconv.Itoa(page)},
			"full_search":  []string{"1"},
			"result_from":  []string{"1"},
			"titleonly":    []string{"3"},
			"replyless":    []string{"0"},
			"replylimit":   []string{"0"},
			"searchdate":   []string{"0"},
			"beforeafter":  []string{"after"},
			"sortby":       []string{"date"},
			"resorder":     []string{"desc"},
			"showposts":    []string{"0"},
			"catlist[]":    []string{"9"},
		},
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

func (is *FrenchStreamReSource) filmSearchPostList(document *crawler.Element) []schema.MoviePost {
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

func (is *FrenchStreamReSource) SerieLatestPostList(page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.SerieLatestURL, page)))
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.serieLatestPostList(crawler.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) serieLatestPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.SerieLatestPostSelector
	serieList := make([]schema.MoviePost, 0)
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
			serieList = append(serieList, schema.MoviePost{
				Category: schema.MovieSerie,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return serieList
}

func (is *FrenchStreamReSource) SerieSearchPostList(query string, page int) []schema.MoviePost {
	response, err := is.PostForm(
		fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.SerieSearchURL, page)),
		url.Values{
			"do":           []string{"search"},
			"subaction":    []string{"search"},
			"story":        []string{query},
			"search_start": []string{strconv.Itoa(page)},
			"full_search":  []string{"1"},
			"result_from":  []string{"1"},
			"titleonly":    []string{"3"},
			"replyless":    []string{"0"},
			"replylimit":   []string{"0"},
			"searchdate":   []string{"0"},
			"beforeafter":  []string{"after"},
			"sortby":       []string{"date"},
			"resorder":     []string{"desc"},
			"showposts":    []string{"0"},
			"catlist[]":    []string{"10"},
		},
	)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.serieSearchPostList(crawler.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) serieSearchPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.SerieSearchPostSelector
	serieList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			serieList = append(serieList, schema.MoviePost{
				Category: schema.MovieSerie,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return serieList
}

func (is *FrenchStreamReSource) FilmArticle(link string) *schema.MovieArticle {
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

func (is *FrenchStreamReSource) filmArticle(document *crawler.Element) *schema.MovieArticle {
	articleSelector := is.FilmArticleSelector

	description := document.ChildText(articleSelector.Description[0])
	// imdb := document.ChildText(articleSelector.Imdb[0])

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(e.ChildText("span"), "Genre") {
				genders = append(genders, e.ChildTexts("a")...)
				return false
			}
			return true
		})

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(e.ChildText("span"), "sortie") {
				date = strings.TrimSpace(e.Selection.Contents().Not("span").Text())
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)

	subtitleHosters := make([]string, 0)
	hosters := make([]string, 0)
	document.ForEach(articleSelector.Hosters[0],
		func(i int, e *crawler.Element) {
			if strings.Contains(strings.ToLower(strings.TrimSpace(e.ChildText("li"))), "vostfr") {
				subtitleHosters = append(subtitleHosters, e.ChildAttribute("li a", "href"))
			}
			if strings.Contains(strings.ToLower(strings.TrimSpace(e.ChildText("li"))), "french") {
				hosters = append(hosters, e.ChildAttribute("li a", "href"))
			}
		})
	videos = append(videos, schema.MovieVideo{
		SubtitleHosters: subtitleHosters,
		Hosters:         hosters,
		Name:            "Film",
	})

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

func (is *FrenchStreamReSource) SerieArticle(link string) *schema.MovieArticle {
	response, err := is.Get(link)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.serieArticle(crawler.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) serieArticle(document *crawler.Element) *schema.MovieArticle {
	articleSelector := is.SerieArticleSelector

	description := document.ChildText(articleSelector.Description[0])
	// imdb := document.ChildText(articleSelector.Imdb[0])

	// var date string
	// document.ForEachWithBreak(articleSelector.Date[0],
	// 	func(i int, e *crawler.Element) bool {
	// 		if strings.Contains(e.ChildText("span"), "sortie") {
	// 			date = strings.TrimSpace(e.Selection.Contents().Not("span").Text())
	// 			return false
	// 		}
	// 		return true
	// 	})

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(e.ChildText("span"), "Genre") {
				genders = append(genders, strings.Split(strings.TrimSpace(e.Selection.Contents().Not("span").Text()), ", ")...)
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)

	videosMap := make(map[string]schema.MovieVideo)
	document.ForEach(articleSelector.Hosters[0],
		func(index int, version *crawler.Element) {
			version.ForEach(articleSelector.Hosters[1], func(i int, episode *crawler.Element) {
				id := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(episode.Attribute("title")), "episode"))
				video := schema.MovieVideo{Name: id, Hosters: make([]string, 0), SubtitleHosters: make([]string, 0)}
				if v, ok := videosMap[id]; ok {
					video = v
				}
				ref := episode.Attribute("data-rel")
				if index == 0 {
					if ref == "" {
						video.Hosters = append(video.Hosters, episode.Attribute("href"))
					} else {
						document.ForEach(fmt.Sprintf("#%v li a", ref), func(i int, hoster *crawler.Element) {
							link := hoster.Attribute("href")
							if link == "" {
								video.Hosters = append(video.Hosters, episode.Attribute("href"))
							} else {
								video.Hosters = append(video.Hosters, link)
							}
						})
					}
				} else {
					if ref == "" {
						video.SubtitleHosters = append(video.SubtitleHosters, episode.Attribute("href"))
					} else {
						document.ForEach(fmt.Sprintf("#%v li a", ref), func(i int, hoster *crawler.Element) {
							link := hoster.Attribute("href")
							if link == "" {
								video.Hosters = append(video.Hosters, episode.Attribute("href"))
							} else {
								video.SubtitleHosters = append(video.SubtitleHosters, link)
							}
						})
					}
				}
				videosMap[id] = video
			})
		})
	for _, v := range videosMap {
		videos = append(videos, v)
	}
	if len(genders) == 0 {
		genders = append(genders, "N/A")
	}
	return &schema.MovieArticle{
		Description: description,
		Genders:     genders,
		Videos:      videos,
		Imdb:        "N/A",
		Date:        "N/A",
	}
}
