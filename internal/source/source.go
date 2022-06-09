package source

import (
	"errors"
	"yola/internal/schema"
)

type Source interface {
	MangaSource
	SerieSource
	FilmSource
}

type MangaSource interface {
	MangaArticle(link string) *schema.MovieArticle
	MangaLatestPostList(page int) []schema.MoviePost
	MangaSearchPostList(query string, page int) []schema.MoviePost
}

type SerieSource interface {
	SerieArticle(link string) *schema.MovieArticle
	SerieLatestPostList(page int) []schema.MoviePost
	SerieSearchPostList(query string, page int) []schema.MoviePost
}

type FilmSource interface {
	FilmArticle(link string) *schema.MovieArticle
	FilmLatestPostList(page int) []schema.MoviePost
	FilmSearchPostList(query string, page int) []schema.MoviePost
}

func ParseMangaSource(name string, source schema.MovieSource) (MangaSource, error) {
	switch name {
	case "french-manga-net":
		return NewFrenchMangaNetSource(&source), nil
	case "vostfree-tv":
		return NewVostFreeSource(&source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func ParseSerieSource(name string, source schema.MovieSource) (SerieSource, error) {
	switch name {
	case "french-stream-re":
		return NewFrenchStreamReSource(&source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func ParseFilmSource(name string, source schema.MovieSource) (FilmSource, error) {
	switch name {
	case "french-stream-re":
		return NewFrenchStreamReSource(&source), nil
	case "illimitestreaming-co":
		return NewIllimitestreamingco(&source), nil
	default:
		return nil, errors.New("no-found")
	}
}
