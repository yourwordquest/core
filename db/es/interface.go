package es

type ElasticSearchDocument interface {
	ESIndex() string
	EsData() (id string, data map[string]interface{})
}
