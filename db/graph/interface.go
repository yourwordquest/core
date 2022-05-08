package graph

type Neo4JUpdater interface {
	GraphUpdate() (entities []GraphEntity)
}
