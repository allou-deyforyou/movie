package firebase

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewFirebaseClient() (*firestore.Client, error) {
	firebaseCredentialsFile := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_FILE_NAME"))
	firebaseProjectID := os.Getenv("FIREBASE_PROJECT_ID")
	client, err := firestore.NewClient(context.Background(), firebaseProjectID, firebaseCredentialsFile)
	if err != nil {
		return nil, err
	}
	return client, nil
}
