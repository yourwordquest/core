package graph

import (
	"fmt"
	"strings"
)

type Node struct {
	Id   string
	Name string
}

type Edge struct {
	Name          string
	Source        string
	SourceId      string
	Destination   string
	DestinationId string
}

type GraphEntity struct {
	*Node
	*Edge
	Data map[string]interface{}
}

type GraphWriter struct {
	entities []GraphEntity
}

func (gw *GraphWriter) AddEntity(entities ...GraphEntity) {
	gw.entities = append(gw.entities, entities...)
}

func (gw *GraphWriter) Write() error {
	transaction, err := session.BeginTransaction()
	if err != nil {
		return err
	}

	// This will rollback the transaction if not committed
	defer transaction.Close()

	for i := range gw.entities {
		entity := gw.entities[i]
		if entity.Data != nil {
			entity.Data = map[string]interface{}{}
		}
		if entity.Node != nil {
			query := nodeQuery(entity.Node, entity.Data)
			entity.Data["__id"] = entity.Node.Id
			_, err := transaction.Run(query, entity.Data)
			if err != nil {
				return err
			}
		}
		if entity.Edge != nil {
			query := edgeQuery(entity.Edge, entity.Data)
			entity.Data["__source"] = entity.Edge.SourceId
			entity.Data["__destination"] = entity.Edge.DestinationId
			_, err := transaction.Run(query, entity.Data)
			if err != nil {
				return err
			}
		}
	}

	return transaction.Commit()
}

func NewWriter(queries []string, params map[string]interface{}) *GraphWriter {
	writer := GraphWriter{
		entities: make([]GraphEntity, 0),
	}
	return &writer
}

func nodeQuery(node *Node, data map[string]interface{}) string {
	params := []string{"Id: $__id"}
	for key := range data {
		params = append(params, fmt.Sprintf("%s $%s", key, key))
	}
	return fmt.Sprintf(`
	MERGE (n:%s {Id: $__id})
	SET n = {%s}
	`, node.Name, strings.Join(params, ", "))
}

func edgeQuery(edge *Edge, data map[string]interface{}) string {
	query := fmt.Sprintf(`
	MATCH
		(source:%s {Id: $__source}),
		(dest:%s {Id: $__destination})
	MERGE (source)-[r:%s]->(dest)
	`, edge.Source, edge.Destination, edge.Name)
	if data != nil {
		params := []string{}
		for key := range data {
			params = append(params, fmt.Sprintf("%s $%s", key, key))
		}
		if len(params) > 0 {
			query = fmt.Sprintf("%s\nSET r = {%s}", query, strings.Join(params, ", "))
		}
	}
	return query
}
