package db

import (
	"github.com/yourwordquest/core/db/elasticsearch"
	"github.com/yourwordquest/core/db/firestore"
	"github.com/yourwordquest/core/db/tigergraph"
)

type MultiDatabaseEntity interface {
	firestore.FirestoreDocument
	elasticsearch.ElasticSearchDocument
	tigergraph.TigerGraphVertex
}

type MultiDatabaseWriter struct {
	fs *firestore.FirestoreWriter
	es *elasticsearch.ElasticSearchWriter
	tg *tigergraph.TigerGraphWriter
}

func (w *MultiDatabaseWriter) AddEntity(entity MultiDatabaseEntity) {
	w.es.AddEntry(entity)
	w.fs.AddDoc(entity)
	w.tg.AddVertex(entity)
}

func (w *MultiDatabaseWriter) AddEdge(edge tigergraph.TigerGraphEdge) {
	w.tg.AddEdge(edge)
}

func (w *MultiDatabaseWriter) Write() (err error) {
	// The order of write is important
	// Firestore writes must succeed first because tiger graph queries return Ids that are referenced from firestore
	// TigerGraph writes must succeed before elasticsearch because tiger graph determines most connections, data must exist there first before it's searched
	err = w.fs.Write()
	if err != nil {
		return
	}
	err = w.tg.Write()
	if err != nil {
		return
	}
	err = w.es.Write()
	return
}

// DeferredWriter is used to write at the end of execution
// This function panics if there's an error on write
func (w *MultiDatabaseWriter) DeferredWrite() {
	err := w.Write()
	if err != nil {
		panic(err)
	}
}

func NewWriter() *MultiDatabaseWriter {
	writer := MultiDatabaseWriter{
		fs: firestore.NewWriter(),
		es: elasticsearch.NewWriter(),
		tg: tigergraph.NewWriter(),
	}
	return &writer
}
