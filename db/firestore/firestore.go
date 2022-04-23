package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func Get(id string, destination FirestoreDocument) (err error) {
	collection := client.Collection(destination.Collection())
	destination.SetId(id)
	doc := collection.Doc(id)
	docSnap, err := doc.Get(context.Background())
	if err != nil {
		return
	}
	err = docSnap.DataTo(destination)
	return
}

type writer struct {
	batch firestore.WriteBatch
}

func (w *writer) AddDoc(doc FirestoreDocument) {
	docRef := client.Doc(fmt.Sprintf("%s/%s", doc.Collection(), doc.Id()))
	w.batch.Set(docRef, doc)
}

func (w *writer) Write() (err error) {
	_, err = w.batch.Commit(context.Background())
	return
}

func Writer() *writer {
	return &writer{
		batch: *client.Batch(),
	}
}
