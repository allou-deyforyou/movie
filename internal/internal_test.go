package internal_test

import (
	"context"
	"log"
	"testing"
	"yola/internal"
	"yola/internal/schema"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

func init() {
	firebaseCredentialsFile := option.WithCredentialsFile("../yola-340622-firebase-adminsdk-823pn-0cad271a6c.json")
	firebaseProjectID := "yola-340622"
	client, err := firestore.NewClient(context.Background(), firebaseProjectID, firebaseCredentialsFile)
	if err != nil {
		panic(err)
	}
	firestoreClient = client
}

func TestCreateIllimitestreamingcoSource(t *testing.T) {
	firestoreClient.Collection(
		schema.MOVIE_SOURCES_COLLECTION,
	).Add(context.Background(), schema.MovieSource{
		Status:        true,
		Name:          "illimitestreaming-co",
		URL:           "https://www.illimitestreaming.co/",
		FilmLatestURL: "/film/page/%v",
		FilmLatestPostSelector: &schema.MoviePostSelector{
			Title: []string{"h2"},
			Image: []string{"img", "data-original"},
			Link:  []string{"a", "href"},
			List:  []string{".movies-list .ml-item"},
		},
		FilmSearchURL: "/page/%v?s=%v",
		FilmSearchPostSelector: &schema.MoviePostSelector{
			Title: []string{"h2"},
			Image: []string{"img", "data-original"},
			Link:  []string{"a", "href"},
			List:  []string{".movies-list .ml-item"},
		},
		FilmArticleSelector: &schema.MovieArticleSelector{
			Imdb:        []string{".mvic-desc p"},
			Genders:     []string{".mvic-desc p"},
			Date:        []string{".mvic-desc p"},
			Hosters:     []string{".idTabs > div"},
			Description: []string{".mvic-desc .desc"},
		},
	})
}

func TestCreateFrenchStreamSource(t *testing.T) {
	firestoreClient.Collection(
		schema.MOVIE_SOURCES_COLLECTION,
	).Add(context.Background(), schema.MovieSource{
		Status:        true,
		Name:          "french-stream-re",
		URL:           "https://french-stream.re",
		FilmLatestURL: "/film/page/%v",
		FilmLatestPostSelector: &schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		},
		FilmSearchURL: "/index.php?do=search",
		FilmSearchPostSelector: &schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		},
		FilmArticleSelector: &schema.MovieArticleSelector{
			Hosters:     []string{"#primary_nav_wrap > ul > li"},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		},

		SerieLatestURL: "/serie/page/%v",
		SerieLatestPostSelector: &schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		},
		SerieSearchURL: "/index.php?do=search",
		SerieSearchPostSelector: &schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		},
		SerieArticleSelector: &schema.MovieArticleSelector{
			Hosters:     []string{".elink", "a"},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		},
	})
}

func TestCreateFrenchMangaNetSource(t *testing.T) {
	firestoreClient.Collection(
		schema.MOVIE_SOURCES_COLLECTION,
	).Add(context.Background(), schema.MovieSource{
		Status:         true,
		Name:           "french-manga-net",
		URL:            "https://french-manga.net",
		MangaLatestURL: "/manga-streaming/page/%v",
		MangaLatestPostSelector: &schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		},
		MangaSearchURL: "/index.php?do=search",
		MangaSearchPostSelector: &schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		},
		MangaArticleSelector: &schema.MovieArticleSelector{
			Hosters:     []string{".elink", "a"},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		},
	})
}

func TestCreateFrenchAnimeSource(t *testing.T) {
	firestoreClient.Collection(
		schema.MOVIE_SOURCES_COLLECTION,
	).Add(context.Background(), schema.MovieSource{
		Status:         true,
		Name:           "french-anime-com",
		URL:            "https://french-anime.com/",
		MangaLatestURL: "/animes-vostfr/page/%v",
		MangaLatestPostSelector: &schema.MoviePostSelector{
			Title: []string{".mov-t"},
			Image: []string{"img", "src"},
			Link:  []string{"a", "href"},
			List:  []string{"#dle-content > .mov"},
		},
		MangaSearchURL: "/index.php?do=search",
		MangaSearchPostSelector: &schema.MoviePostSelector{
			Title: []string{".mov-t"},
			Image: []string{"img", "src"},
			Link:  []string{"a", "href"},
			List:  []string{"#dle-content > .mov"},
		},
		MangaArticleSelector: &schema.MovieArticleSelector{
			Hosters:     []string{".eps"},
			Imdb:        []string{""},
			Description: []string{".mov-list li", ".mov-desc"},
			Genders:     []string{".mov-list li", ".mov-desc"},
			Date:        []string{".mov-list li", ".mov-desc"},
		},
	})
}

func TestGetAllMovieService(t *testing.T) {
	sources := internal.GetAllMovieService(
		firestoreClient.Collection(
			schema.MOVIE_SOURCES_COLLECTION,
		).Where("status", "==", true),
	)
	log.Println(sources)
}
