package source

import (
	"net/http"
	"server/lib/schema"
)

type Source interface {
	SearchFilmPostList(query string, page int) []schema.MoviePost
	SetData(source schema.MovieSource, client *http.Client)
	FilmArticle(link string) *schema.MovieFilmArticle
	FilmPostList(page int) []schema.MoviePost
	Verify(string) bool
}

func GetAllSources() []Source {
	return []Source{&IllimitestreamingcoSource{}, &Vffilmtop{}}
}
