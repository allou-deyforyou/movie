package main

import (
	"context"
	"strings"

	"yola/internal"
	"yola/internal/schema"
	"yola/internal/source"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var sources = make([]schema.MovieSource, 0)

func init() {
	firestoreClient, err := internal.NewFirebaseClient()
	if err != nil {
		panic(err)
	}
	sources = internal.GetAll[schema.MovieSource](
		firestoreClient.Collection(
			schema.MOVIE_SOURCES_COLLECTION,
		).Where("status", "==", true),
	)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters["source"]
	link := request.QueryStringParameters["link"]
	response := new(schema.MovieArticle)
	for _, movieSource := range sources {
		if strings.EqualFold(name, movieSource.Name) {
			if source, err := source.ParseMangaSource(name, movieSource); err == nil {
				response = source.MangaArticle(link)
			}
			break
		}
	}
	return internal.ServeResponse(response)
}
