package main

import (
	"context"
	"net/http"
	"server/lib/firebase"
	"server/lib/handler"
	"server/lib/schema"
	"server/lib/source"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var allSources []source.Source

func init() {
	allSources = source.GetAllSources()
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (handler.Response, error) {
	name := request.QueryStringParameters["source"]
	link := request.QueryStringParameters["link"]

	firestoreClient, err := firebase.NewFirebaseClient()
	if err != nil {
		panic(err)
	}
	collection := firestoreClient.Collection(schema.MOVIE_FILM_SOURCES_COLLECTION)
	query := collection.Where("name", "==", name)
	movieSource := firebase.GetAll[schema.MovieFilmSource](query)[0]
	var article *schema.MovieFilmArticle
	for _, source := range allSources {
		if source.Verify(movieSource.Name) {
			source.SetData(movieSource, http.DefaultClient)
			article = source.FilmArticle(link)
		}
	}
	return handler.ServeResponse(article)
}

func main() {
	lambda.Start(Handler)
}
