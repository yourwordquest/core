package promise

import (
	"github.com/yourwordquest/core/common"
	"github.com/yourwordquest/core/db"
	"github.com/yourwordquest/core/db/es"
	"github.com/yourwordquest/core/utils"
)

type Promise struct {
	id                string
	Summary           string           `firestore:"Summary"`
	Details           string           `firestore:"Details"`
	PromisedOn        int64            `firestore:"PromisedOn"`
	ExpectedDelivery  int64            `firestore:"ExpectedDelivery"`
	DeliveredOn       int64            `firestore:"DeliveredOn"`
	AddedOn           int64            `firestore:"AddedOn"`
	SupportingContent string           `firestore:"SupportingContent"`
	Status            string           `firestore:"Status"`
	OtherData         common.OtherData `firestore:"OtherData"`
}

func (prom *Promise) Id() string {
	return prom.id
}

func (prom *Promise) SetId(id string) {
	prom.id = id
}

func (prom *Promise) Collection() string {
	return "Promises"
}

func (prom *Promise) ESIndex() string {
	return "searchable_promises"
}

func (prom *Promise) EsData() (id string, data map[string]interface{}) {
	id = prom.id
	data = map[string]interface{}{
		"Summary":          prom.Summary,
		"Detail":           utils.StripHTML(prom.Details),
		"PromisedOn":       prom.PromisedOn,
		"ExpectedDelivery": prom.ExpectedDelivery,
		"DeliveredOn":      prom.DeliveredOn,
		"AddedOn":          prom.AddedOn,
		"Status":           prom.Status,
	}
	return
}

var _ db.FirestoreDocument = new(Promise)
var _ es.ElasticSearchDocument = new(Promise)
