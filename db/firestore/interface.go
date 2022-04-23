package firestore

type FirestoreDocument interface {
	Id() string
	SetId(id string)
	Collection() string
}
