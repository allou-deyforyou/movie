package source

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"yola/internal/crawler"
	"yola/internal/schema"

	"github.com/PuerkitoBio/goquery"
)

type VostFreeSource struct {
	*schema.MovieSource
	*http.Client
}

func NewVostFreeSource(source *schema.MovieSource) *VostFreeSource {
	return &VostFreeSource{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *VostFreeSource) MangaLatestPostList(page int) []schema.MoviePost {
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

func (is *VostFreeSource) mangaLatestPostList(document *crawler.Element) []schema.MoviePost {
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

func (is *VostFreeSource) MangaSearchPostList(query string, page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.MangaSearchURL, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.mangaSearchPostList(crawler.NewElement(document.Selection))
}

func (is *VostFreeSource) mangaSearchPostList(document *crawler.Element) []schema.MoviePost {
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

func (is *VostFreeSource) MangaArticle(link string) *schema.MovieArticle {
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

func (is *VostFreeSource) mangaArticle(document *crawler.Element) *schema.MovieArticle {
	articleSelector := is.MangaArticleSelector
	description := strings.Join(strings.Fields(document.ChildText(articleSelector.Description[0])), " ")
	// imdb := document.ChildText(articleSelector.Imdb[0])

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("span")), "année") {
				date = strings.TrimSpace(e.ChildText("a"))
				return false
			}
			return true
		})

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *crawler.Element) bool {
			if strings.Contains(strings.ToLower(e.Text()), "genre") {
				genders = append(genders, e.ChildTexts("a")...)
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)
	for _, videoSelector := range articleSelector.Videos {
		document.ForEach(videoSelector.Hosters[0],
			func(index int, episode *crawler.Element) {
				subtitleHosters := make([]string, 0)
				hosters := make([]string, 0)
				episode.ForEach("div",
					func(_ int, hoster *crawler.Element) {
						id := fmt.Sprintf("#content_%v", hoster.Attribute("id"))
						link := document.ChildText(id)
						switch strings.ToLower(hoster.Text()) {
						case "uqload":
							link = fmt.Sprintf("https://uqload.com/embed-%v.html", link)
						case "sibnet":
							link = fmt.Sprintf("https://video.sibnet.ru/shell.php?videoid=%v", link)
						case "mytv":
							link = fmt.Sprintf("https://www.myvi.tv/embed/%v", link)
						}
						if strings.Contains(strings.ToLower(document.ChildText(".slide-middle h1")), "vf") {
							hosters = append(hosters, link)
						} else {
							subtitleHosters = append(subtitleHosters, link)
						}
					})
				videos = append(videos, schema.MovieVideo{
					Name:            strconv.Itoa(index + 1),
					SubtitleHosters: subtitleHosters,
					Hosters:         hosters,
				})
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
