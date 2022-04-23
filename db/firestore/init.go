package firestore

import (
	"context"
	"os"

	firestore "cloud.google.com/go/firestore"
)

var client *firestore.Client

func init() {
	project_id := os.Getenv("GOOGLE_PROJECT_ID")
	_client, err := firestore.NewClient(context.Background(), project_id)
	if err != nil {
		panic(err)
	}

	client = _client
}
