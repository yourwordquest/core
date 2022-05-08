package graph

type Neo4JDocument interface {
	GraphUpdate() (entities []GraphEntity)
}
