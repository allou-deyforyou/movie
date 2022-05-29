package schema

const (
	MOVIE_SERIETV_SOURCES_COLLECTION = "movie_serietv_sources"
	MOVIE_FILM_SOURCES_COLLECTION    = "movie_film_sources"
)

type MovieFilmSource struct {
	MovieFilmArticleSelector MovieFilmArticleSelector `json:"article_selector"`
	LatestFilmsURL           string                   `json:"latest_films_url"`
	SearchFilmsURL           string                   `json:"search_films_url"`
	FilmPostSelector         MoviePostSelector        `json:"post_selector"`
	Status                   bool                     `json:"status"`
	Name                     string                   `json:"name"`
	URL                      string                   `json:"url"`
}
