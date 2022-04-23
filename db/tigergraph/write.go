package tigergraph

import (
	"errors"
	"fmt"

	"github.com/yourwordquest/core/utils/requests"
)

type TigerGraphWriter struct {
	vertices    map[string]vertexObject
	edges       map[string]edgeObject
	vertexCount int
	edgeCount   int
}

func (w *TigerGraphWriter) AddVertex(vertices ...TigerGraphVertex) *TigerGraphWriter {
	w.vertexCount += len(vertices)
	for i := range vertices {
		vertex := vertices[i]
		vertexTypeData, ok := w.vertices[vertex.TgVertex()]
		if !ok {
			vertexTypeData = vertexObject{}
			w.vertices[vertex.TgVertex()] = vertexTypeData
		}

		id, data := vertex.TgData()
		vertexTypeData[id] = data
	}
	return w
}

func (w *TigerGraphWriter) AddEdge(edges ...TigerGraphEdge) *TigerGraphWriter {
	w.edgeCount += len(edges)
	for i := range edges {
		edge := edges[i]
		sourceVertex, ok := w.edges[edge.SourceVertex()]
		if !ok {
			sourceVertex = edgeObject{}
			w.edges[edge.SourceVertex()] = sourceVertex
		}
		source, target, data := edge.TgData()
		sourceDetails, ok := sourceVertex[source]
		if !ok {
			sourceDetails = map[string]map[string]map[string]TigerGraphObject{}
			sourceVertex[source] = sourceDetails
		}

		edgeTypeDetails, ok := sourceDetails[edge.EdgeType()]
		if !ok {
			edgeTypeDetails = map[string]map[string]TigerGraphObject{}
			sourceDetails[edge.EdgeType()] = edgeTypeDetails
		}

		targetDetails, ok := edgeTypeDetails[edge.TargetVertex()]
		if !ok {
			targetDetails = map[string]TigerGraphObject{}
			edgeTypeDetails[edge.EdgeType()] = targetDetails
		}

		targetDetails[target] = data
	}
	return w
}

// Write pushes all changes to TigerGraph, the changes will be upserted atomically
func (w *TigerGraphWriter) Write() (err error) {
	if w.edgeCount == 0 && w.vertexCount == 0 {
		return nil
	}

	token, err := GetToken()
	if err != nil {
		return
	}

	payload := requests.Payload{}
	if w.edgeCount > 0 {
		payload["edges"] = w.edges
	}
	if w.vertexCount > 0 {
		payload["vertices"] = w.vertices
	}

	url := fmt.Sprintf("%s:9000/graph/%s", server, graph)
	resp, err := requests.Request{
		JSON: payload,
		Headers: requests.Headers{
			"gsql-atomic-level": "atomic",
			"Authorization":     fmt.Sprintf("Bearer %s", token),
		},
	}.Send(url, nil)

	if 200 < resp.StatusCode || resp.StatusCode > 300 {
		err = errors.New(string(resp.Body))
	}
	return
}

// DeferredWriter is used to write at the end of execution
// This function panics if there's an error on write
func (w *TigerGraphWriter) DeferredWrite() {
	err := w.Write()
	if err != nil {
		panic(err)
	}
}

func NewWriter() *TigerGraphWriter {
	return &TigerGraphWriter{
		vertices:    make(map[string]vertexObject),
		vertexCount: 0,
		edges:       make(map[string]edgeObject),
		edgeCount:   0,
	}
}
