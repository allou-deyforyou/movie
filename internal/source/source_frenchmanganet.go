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

type FrenchMangaNetSource struct {
	*schema.MovieSource
	*http.Client
}

func NewFrenchMangaNetSource(source *schema.MovieSource) *FrenchMangaNetSource {
	return &FrenchMangaNetSource{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *FrenchMangaNetSource) MangaLatestPostList(page int) []schema.MoviePost {
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

func (is *FrenchMangaNetSource) mangaLatestPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.MangaLatestPostSelector
	mangaList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
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

func (is *FrenchMangaNetSource) MangaSearchPostList(query string, page int) []schema.MoviePost {
	response, err := is.PostForm(
		fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.MangaSearchURL, page)),
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
			"catlist[]":    []string{"16"},
		},
	)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.mangaSearchPostList(crawler.NewElement(document.Selection))
}

func (is *FrenchMangaNetSource) mangaSearchPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.MangaSearchPostSelector
	mangaList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])

			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			mangaList = append(mangaList, schema.MoviePost{
				Category: schema.MovieManga,
				Source: is.Name,
				Image:  image,
				Title:  title,
				Link:   link,
			})
		})
	return mangaList
}

func (is *FrenchMangaNetSource) MangaArticle(link string) *schema.MovieArticle {
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

func (is *FrenchMangaNetSource) mangaArticle(document *crawler.Element) *schema.MovieArticle {
	articleSelector := is.MangaArticleSelector
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
				genders = append(genders, strings.Split(strings.TrimSpace(e.Selection.Contents().Not("span").Text()), " - ")...)
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)
	for _, videoSelector := range articleSelector.Videos {
		videosMap := make(map[string]schema.MovieVideo)
		document.ForEach(videoSelector.Hosters[0],
			func(index int, version *crawler.Element) {
				version.ForEach(videoSelector.Hosters[1], func(i int, episode *crawler.Element) {
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
								link := strings.TrimSpace(hoster.Attribute("href"))
								if link == "" {
									video.Hosters = append(video.Hosters, episode.Attribute("href"))
								} else {
									if hoster.ChildAttribute("i", "aria-hidden") != "" {
										video.Hosters = append(video.Hosters, link)
									}
								}
							})
						}
					} else {
						if ref == "" {
							video.SubtitleHosters = append(video.SubtitleHosters, episode.Attribute("href"))
						} else {
							document.ForEach(fmt.Sprintf("#%v li a", ref), func(i int, hoster *crawler.Element) {
								link := strings.TrimSpace(hoster.Attribute("href"))
								if link == "" {
									video.Hosters = append(video.Hosters, episode.Attribute("href"))
								} else {
									if hoster.ChildAttribute("i", "aria-hidden") != "" {
										video.SubtitleHosters = append(video.SubtitleHosters, link)
									}
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
