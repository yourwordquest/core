package db

type FirestoreDocument interface {
	Id() string
	SetId(id string)
	Collection() string
}
