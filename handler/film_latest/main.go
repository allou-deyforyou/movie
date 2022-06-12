package main

import (
	"context"
	"log"
	"math/rand"
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

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	context, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(context,
		chromedp.Navigate(`https://vostfree.tv`),
		chromedp.Text(`body`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))

	var movieSources []source.FilmSource
	for _, movieSource := range sources {
		if source, err := source.ParseFilmSource(movieSource.Name, movieSource); err == nil {
			movieSources = append(movieSources, source)
		}
	}
	page, _ := strconv.Atoi(request.PathParameters["page"])
	moviePosts := make([]schema.MoviePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range movieSources {
		group.Add(1)
		go func(source source.FilmSource) {
			posts := source.FilmLatestPostList(page)
			moviePosts = append(moviePosts, posts...)
			group.Done()
		}(s)
	}
	group.Wait()
	length := len(moviePosts)
	response := make([]schema.MoviePost, length)
	perm := rand.Perm(length)
	for i, v := range perm {
		response[v] = moviePosts[i]
	}

	return internal.ServeResponse(response)
}
