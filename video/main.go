package main

import (
	"context"
	"log"
	"net/http"
	"server/lib/firebase"
	"server/lib/handler"
	"server/lib/hoster"
	"server/lib/schema"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var myHoster hoster.Hoster

func init() {
	allHosters := hoster.GetAllHosters()
	firestoreClient, err := firebase.NewFirebaseClient()
	if err != nil {
		panic(err)
	}

	collection := firestoreClient.Collection(schema.MOVIE_HOSTERS_COLLECTION)
	query := collection.Where("status", "==", true)
	movieSources := firebase.GetAll[schema.MovieHoster](query)
	for _, moviehoster := range movieSources {
		for _, h := range allHosters {
			if h.Verify(moviehoster.Name) {
				h.SetData(moviehoster, http.DefaultClient)
				myHoster = h
			}
		}
	}
	log.Println(myHoster)
	// go firebase.Snapshots(firestoreClient.Collection(schema.MOVIE_HOSTERS_COLLECTION), func(data []schema.MovieHoster) {
	// 	for _, moviehoster := range data {
	// 		for _, h := range allHosters {
	// 			if h.Verify(moviehoster.Name) {
	// 				h.SetData(moviehoster, http.DefaultClient)
	// 				myHoster = h
	// 			}
	// 		}
	// 	}
	// })
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (handler.Response, error) {
	referer := request.QueryStringParameters["referer"]
	link := request.QueryStringParameters["link"]
	result := myHoster.Video(link, referer)
	return handler.ServeResponse(result)
}

func main() {
	lambda.Start(Handler)
}
