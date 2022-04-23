package db

type FirestoreDocumnet interface {
	Id() string
	SetId(id string)
	Collection() string
}

type ElasticSearchDocument interface {
	ESIndex() string
	EsData() (id string, data map[string]interface{})
}
