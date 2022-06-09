package internal_test

import (
	"context"
	"testing"
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
			Videos: []schema.MovieVideoSelector{
				{Hosters: []string{"#primary_nav_wrap > ul > li"}},
			},
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
			Videos: []schema.MovieVideoSelector{
				{Hosters: []string{".elink", "a"}},
			},
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
			Videos: []schema.MovieVideoSelector{
				{Hosters: []string{".elink", "a"}},
			},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		},
	})
}

func TestCreateNekoSamaFrSource(t *testing.T) {
	firestoreClient.Collection(
		schema.MOVIE_SOURCES_COLLECTION,
	).Add(context.Background(), schema.MovieSource{
		Status:         true,
		Name:           "vostfree-tv",
		URL:            "https://vostfree.tv",
		MangaLatestURL: "/animes-vostfr/page/%v",
		MangaLatestPostSelector: &schema.MoviePostSelector{
			Title: []string{".title"},
			Image: []string{"img", "src"},
			Link:  []string{"a", "href"},
			List:  []string{"#dle-content .movie-poster"},
		},
		MangaSearchURL: "/",
		MangaSearchPostSelector: &schema.MoviePostSelector{
			Title: []string{".title"},
			Image: []string{"img", "src"},
			Link:  []string{"a", "href"},
			List:  []string{"#dle-content .search-result"},
		},
		MangaArticleSelector: &schema.MovieArticleSelector{
			Videos: []schema.MovieVideoSelector{
				{Hosters: []string{".new_player_bottom .button_box"}},
			},
			Imdb:        []string{""},
			Description: []string{".slide-middle .slide-desc"},
			Genders:     []string{".slide-middle .right"},
			Date:        []string{".slide-info p"},
		},
	})
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
			Videos: []schema.MovieVideoSelector{
				{Hosters: []string{".idTabs > #list-downloads"}},
			},
			Imdb:        []string{".mvic-desc p"},
			Genders:     []string{".mvic-desc p"},
			Date:        []string{".mvic-desc p"},
			Description: []string{".mvic-desc .desc"},
		},
	})

}
