package tigergraph

type TigerGraphObject map[string]interface{}

type TigerGraphVertex interface {
	TgVertex() string
	TgData() (id string, data TigerGraphObject)
}

type TigerGraphEdge interface {
	SourceVertex() string
	TargetVertex() string
	EdgeType() string
	TgData() (source string, target string, data TigerGraphObject)
}

type edgeObject map[string]map[string]map[string]map[string]TigerGraphObject
type vertexObject map[string]TigerGraphObject
