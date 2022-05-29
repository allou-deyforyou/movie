package main

import (
	"context"
	"net/http"
	"strconv"

	"server/lib/firebase"
	"server/lib/handler"
	"server/lib/schema"
	"server/lib/source"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (handler.Response, error) {
	query := request.PathParameters["query"]
	page, err := strconv.Atoi(request.PathParameters["page"])
	if err != nil {
		page = 1
	}
	result := make([]schema.MoviePost, 0)
	for _, source := range sources {
		result = append(result, source.SearchFilmPostList(query, page)...)
	}
	return handler.ServeResponse(result)
}

func main() {
	lambda.Start(Handler)
}
