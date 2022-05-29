package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewFirebaseClient() (*firestore.Client, error) {
	firebaseCredentialsFile := option.WithCredentialsFile("yola-340622-firebase-adminsdk-823pn-0cad271a6c.json") // os.Getenv("FIREBASE_CREDENTIALS_FILE_NAME"))
	firebaseProjectID := "yola-340622"                                                                           // os.Getenv("FIREBASE_PROJECT_ID")
	client, err := firestore.NewClient(context.Background(), firebaseProjectID, firebaseCredentialsFile)
	if err != nil {
		return nil, err
	}
	return client, nil
}
