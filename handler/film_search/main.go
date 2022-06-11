package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"

	"yola/internal"
	"yola/internal/schema"
	"yola/internal/source"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chromedp/chromedp"
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
	go func(sources *[]schema.MovieSource) {
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
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://vostfree.tv`),
		chromedp.InnerHTML(`body`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(strings.TrimSpace(res))

	//https://uqload.com/embed-f0adetjw8mhp.html

	var movieSources []source.FilmSource
	for _, movieSource := range sources {
		if source, err := source.ParseFilmSource(movieSource.Name, movieSource); err == nil {
			movieSources = append(movieSources, source)
		}
	}
	query := request.QueryStringParameters["query"]
	page, _ := strconv.Atoi(request.PathParameters["page"])
	response := make([]schema.MoviePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range movieSources {
		group.Add(1)
		go func(source source.FilmSource) {
			posts := source.FilmSearchPostList(query, page)
			response = append(response, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	return internal.ServeResponse(response)
}
