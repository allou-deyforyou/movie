package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"yola/internal/schema"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"github.com/aws/aws-lambda-go/events"
)

func NewFirebaseClient() (*firestore.Client, error) {
	firebaseCredentialsFile := option.WithCredentialsFile("yola-340622-firebase-adminsdk-823pn-0cad271a6c.json")
	firebaseProjectID := "yola-340622"
	client, err := firestore.NewClient(context.Background(), firebaseProjectID, firebaseCredentialsFile)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAllMovieService(collection firestore.Query) []schema.MovieSource {
	documentIterator := collection.Documents(context.Background())
	documentList, err := documentIterator.GetAll()
	if err != nil {
		panic(err)
	}
	results := make([]schema.MovieSource, 0)
	for _, document := range documentList {
		data := new(schema.MovieSource)
		b, _ := json.Marshal(document.Data())
		json.Unmarshal(b, data)
		results = append(results, *data)
	}
	return results
}

func ServeResponse(data interface{}) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
	}
	return resp, nil
}
