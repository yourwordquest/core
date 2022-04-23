package db

type TigerGraphVertex interface {
	TgVertex() string
	TgData() (id string, data map[string]interface{})
}

type TigerGraphEdge interface {
	SourceVertex() string
	TargetVertex() string
	TgData() (source string, target string, data map[string]interface{})
}

type FirestoreDocumnet interface {
	Id() string
	SetId(id string)
	Collection() string
}

type ElasticSearchDocument interface {
	ESIndex() string
	EsData() (id string, data map[string]interface{})
}
