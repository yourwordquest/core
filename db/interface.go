package db

type FirestoreDocumnet interface {
	Id() string
	SetId(id string)
	Collection() string
}
