package schema

const (
	MOVIE_SOURCES_COLLECTION = "movie_sources"
)

type MovieSource struct {
	MangaLatestPostSelector *MoviePostSelector    `firestore:"manga_latest_post_selector,omitempty" json:"manga_latest_post_selector,omitempty"`
	MangaLatestURL          string                `firestore:"manga_latest_url,omitempty" json:"manga_latest_url,omitempty"`
	MangaSearchPostSelector *MoviePostSelector    `firestore:"manga_search_post_selector,omitempty" json:"manga_search_post_selector,omitempty"`
	MangaSearchURL          string                `firestore:"manga_search_url,omitempty" json:"manga_search_url,omitempty"`
	MangaArticleSelector    *MovieArticleSelector `firestore:"manga_article_selector,omitempty" json:"manga_article_selector,omitempty"`

	SerieLatestPostSelector *MoviePostSelector    `firestore:"serie_latest_post_selector,omitempty" json:"serie_latest_post_selector,omitempty"`
	SerieLatestURL          string                `firestore:"serie_latest_url,omitempty" json:"serie_latest_url,omitempty"`
	SerieSearchPostSelector *MoviePostSelector    `firestore:"serie_search_post_selector,omitempty" json:"serie_search_post_selector,omitempty"`
	SerieSearchURL          string                `firestore:"serie_search_url,omitempty" json:"serie_search_url,omitempty"`
	SerieArticleSelector    *MovieArticleSelector `firestore:"serie_article_selector,omitempty" json:"serie_article_selector,omitempty"`

	FilmLatestPostSelector *MoviePostSelector    `firestore:"film_latest_post_selector,omitempty" json:"film_latest_post_selector,omitempty"`
	FilmLatestURL          string                `firestore:"film_latest_url,omitempty" json:"film_latest_url,omitempty"`
	FilmSearchPostSelector *MoviePostSelector    `firestore:"film_search_post_selector,omitempty" json:"film_search_post_selector,omitempty"`
	FilmSearchURL          string                `firestore:"film_search_url,omitempty" json:"film_search_url,omitempty"`
	FilmArticleSelector    *MovieArticleSelector `firestore:"film_article_selector,omitempty" json:"film_article_selector,omitempty"`

	Status bool   `firestore:"status,omitempty" json:"status,omitempty"`
	Name   string `firestore:"name,omitempty" json:"name,omitempty"`
	URL    string `firestore:"url,omitempty" json:"url,omitempty"`
}
