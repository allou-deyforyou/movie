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

var sources = make([]schema.MovieSource, 0)

func init() {
	firestoreClient, err := internal.NewFirebaseClient()
	if err != nil {
		panic(err)
	}
	sources = internal.GetAllMovieService(
		firestoreClient.Collection(
			schema.MOVIE_SOURCES_COLLECTION,
		).Where("status", "==", true),
	)
	go func (sources *[]schema.MovieSource)  {
		internal.SnapshotAllMovieService(
		firestoreClient.Collection(
			schema.MOVIE_SOURCES_COLLECTION,
		).Where("status", "==", true),
		func(ms []schema.MovieSource) {
			*sources = ms
		},
	)
	}(&sources)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movieSources []source.MangaSource
	for _, movieSource := range sources {
		if source, err := source.ParseMangaSource(movieSource.Name, movieSource); err == nil {
			movieSources = append(movieSources, source)
		}
	}
	query := request.QueryStringParameters["query"]
	page, _ := strconv.Atoi(request.PathParameters["page"])
	response := make([]schema.MoviePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range movieSources {
		group.Add(1)
		go func(source source.MangaSource) {
			posts := source.MangaSearchPostList(query, page)
			response = append(response, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	return internal.ServeResponse(response)
}
