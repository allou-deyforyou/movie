package schema

type MovieFilmArticle struct {
	Description string   `json:"description"`
	Hosters     []string `json:"hosters"`
	Genders     []string `json:"genders"`
	Date        string   `json:"date"`
	Imdb        string   `json:"imdb"`
}

type MovieFilmArticleSelector struct {
	Description []string `json:"description"`
	Hosters     []string `json:"hosters"`
	Genders     []string `json:"genders"`
	Date        []string `json:"date"`
	Imdb        []string `json:"imdb"`
}
