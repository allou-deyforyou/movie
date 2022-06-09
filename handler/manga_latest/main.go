package main

import (
	"context"
	"strconv"
	"sync"

	"yola/internal"
	"yola/internal/schema"
	"yola/internal/source"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var sources = make([]source.MangaSource, 0)

func init() {
	firestoreClient, err := internal.NewFirebaseClient()
	if err != nil {
		panic(err)
	}
	movieSources := internal.GetAllMovieService(
		firestoreClient.Collection(
			schema.MOVIE_SOURCES_COLLECTION,
		).Where("status", "==", true),
	)
	for _, movieSource := range movieSources {
		if source, err := source.ParseMangaSource(movieSource.Name, movieSource); err == nil {
			sources = append(sources, source)
		}
	}
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	page, _ := strconv.Atoi(request.PathParameters["page"])
	response := make([]schema.MoviePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range sources {
		group.Add(1)
		go func(source source.MangaSource) {
			posts := source.MangaLatestPostList(page)
			response = append(response, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	return internal.ServeResponse(response)
}
