package source

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"server/lib/crawler"
	"server/lib/schema"

	"github.com/PuerkitoBio/goquery"
)

type Mesfilmsnet struct {
	schema.MovieFilmSource
	*http.Client
}

func (is *Mesfilmsnet) Verify(name string) bool {
	return strings.Contains("mesfilmsnet", name)
}

func (is *Mesfilmsnet) SetData(source schema.MovieFilmSource, client *http.Client) {
	is.MovieFilmSource = source
	is.Client = client
}

func (is *Mesfilmsnet) FilmPostList(page int) []schema.MoviePost {
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

func (is *Mesfilmsnet) filmPostList(document *crawler.Element) []schema.MoviePost {
	selector := is.FilmPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *crawler.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			filmList = append(filmList, schema.MoviePost{
				Image:  strings.ReplaceAll(image, ".jpg", "h.jpg"),
				Source: is.Name,
				Title:  title,
				Link:   link,
			})
		})
	return filmList
}

func (is *Mesfilmsnet) SearchFilmPostList(query string, page int) []schema.MoviePost {
	response, err := is.Get(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(is.LatestFilmsURL, query, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmPostList(crawler.NewElement(document.Selection))
}

func (is *Mesfilmsnet) FilmArticle(link string) *schema.MovieFilmArticle {
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

func (is *Mesfilmsnet) filmArticle(document *crawler.Element) *schema.MovieFilmArticle {
	articleSelector := is.MovieFilmArticleSelector
	posts := document.ChildAttributes(articleSelector.Hosters[0], "data-post")
	numes := document.ChildAttributes(articleSelector.Hosters[0], "data-nume")
	var hosters []string
	for i, post := range posts {
		nume := numes[i]
		response, _ := http.PostForm(fmt.Sprintf("%s/wp-admin/admin-ajax.php", is.URL), url.Values{
			"action": []string{"doo_player_ajax"},
			"type":   []string{"movie"},
			"post":   []string{post},
			"nume":   []string{nume},
		})
		b, _ := io.ReadAll(response.Body)
		var data map[string]string
		json.Unmarshal(b, &data)
		hosters = append(hosters, data["embed_url"])
	}
	return &schema.MovieFilmArticle{Hosters: hosters}
}
