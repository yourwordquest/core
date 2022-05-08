package graph

type GraphData map[string]interface{}

type Neo4JNode interface {
	Node() string
	GraphData() (id string, data GraphData)
	GraphEdges() []Neo4JEdge
}

type Neo4JEdge interface {
	SourceNode() string
	TargetNode() string
	EdgeType() string
	EdgeData(source string, target string, data GraphData)
}
