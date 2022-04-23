package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/yourwordquest/core/utils/requests"
)

type ElasticSearchWriter struct {
	entries []string
}

func (w *ElasticSearchWriter) AddEntry(doc ElasticSearchDocument) *ElasticSearchWriter {
	// Add a bul entry
	id, data := doc.EsData()
	indexBytes, _ := json.Marshal(map[string]map[string]string{
		"index": {
			"_index": doc.ESIndex(),
			"_id":    id,
		},
	})
	dataBytes, _ := json.Marshal(data)

	w.entries = append(w.entries, string(indexBytes), string(dataBytes))
	return w
}

func (w *ElasticSearchWriter) Write() (err error) {
	// Last entry must be a new line
	w.entries = append(w.entries, "")
	payload := strings.Join(w.entries, "\n")
	resp, err := requests.Request{
		Headers: requests.Headers{
			"Authorization": fmt.Sprintf("Basic %s:%s", username, password),
			"Content-Type":  "application/x-ndjson",
		},
		Payload: []byte(payload),
	}.Send(fmt.Sprintf("%s/_bulk", server), nil)

	if 200 < resp.StatusCode || resp.StatusCode > 300 {
		err = errors.New(string(resp.Body))
	}
	return
}

// DeferredWriter is used to write at the end of execution
// This function panics if there's an error on write
func (w *ElasticSearchWriter) DeferredWrite() {
	err := w.Write()
	if err != nil {
		panic(err)
	}
}

func NewWriter() *ElasticSearchWriter {
	return &ElasticSearchWriter{
		entries: make([]string, 0),
	}
}
