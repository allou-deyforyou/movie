package schema

type MovieFilmArticle struct {
	Description string   `json:"description"`
	Hosters     []string `json:"hosters"`
}

type MovieSerieArticle struct {
}

type MovieSeasonArticle struct {
}

type MovieArticleResponse[T interface{}] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type MovieArticleReqest struct {
	Type string `json:"type"`
	Link string `json:"link"`
}
