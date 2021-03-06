package db

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

type FirestoreWriter struct {
	batch firestore.WriteBatch
}

func (w *FirestoreWriter) AddDoc(doc FirestoreDocument) {
	docRef := client.Doc(fmt.Sprintf("%s/%s", doc.Collection(), doc.Id()))
	w.batch.Set(docRef, doc)
}

func (w *FirestoreWriter) Write() (err error) {
	_, err = w.batch.Commit(context.Background())
	return
}

// DeferredWriter is used to write at the end of execution
// This function panics if there's an error on write
func (w *FirestoreWriter) DeferredWrite() {
	err := w.Write()
	if err != nil {
		panic(err)
	}
}

func NewWriter() *FirestoreWriter {
	return &FirestoreWriter{
		batch: *client.Batch(),
	}
}
