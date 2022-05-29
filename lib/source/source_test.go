package source_test

import (
	"log"
	"net/http"
	"server/lib/firebase"
	"server/lib/schema"
	"server/lib/source"
	"testing"
)

var sources []source.Source

func init() {
	allSources := source.GetAllSources()

	firestoreClient, err := firebase.NewFirebaseClient()
	if err != nil {
		panic(err)
	}
	collection := firestoreClient.Collection(schema.MOVIE_FILM_SOURCES_COLLECTION)
	query := collection.Where("status", "==", true)
	sources = parseSourceList(firebase.GetAll[schema.MovieFilmSource](query), allSources)
}

func parseSourceList(data []schema.MovieFilmSource, allSources []source.Source) []source.Source {
	var result []source.Source
	for _, movieSource := range data {
		for _, source := range allSources {
			if source.Verify(movieSource.Name) {
				source.SetData(movieSource, http.DefaultClient)
				result = append(result, source)
			}
		}
	}
	return result
}

func Test(t *testing.T) {
	for _, source := range sources {
		log.Println(source.SearchFilmPostList("jumanji", 1))
	}
}
