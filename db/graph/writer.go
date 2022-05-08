package graph

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type GraphWriter struct {
	transaction neo4j.Transaction
}

func (gw *GraphWriter) AddNode(node Neo4JNode) {

}

func (gw *GraphWriter) Write() (err error) {
	return nil
}

func NewWriter() *GraphWriter {
	transaction, _ := session.BeginTransaction()

	writer := GraphWriter{
		transaction: transaction,
	}
	return &writer
}
