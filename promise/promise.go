package promise

import (
	"github.com/yourwordquest/core/common"
	"github.com/yourwordquest/core/db"
	"github.com/yourwordquest/core/db/tigergraph"
	"github.com/yourwordquest/core/utils"
)

type Promise struct {
	id                string
	Summary           string
	Details           string
	PromisedOn        int64
	ExpectedDelivery  int64
	DeliveredOn       int64
	AddedOn           int64
	SupportingContent string
	Status            string
	OtherData         common.OtherData
}

func (prom *Promise) Id() string {
	return prom.id
}

func (prom *Promise) SetId(id string) {
	prom.id = id
}

func (prom *Promise) Collection() string {
	return "promises"
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

func (prom *Promise) TgVertex() string {
	return "Promise"
}

func (prom *Promise) TgData() (id string, data tigergraph.TigerGraphObject) {
	id = prom.id
	data = tigergraph.TigerGraphObject{
		"PromisedOn":       prom.PromisedOn,
		"ExpectedDelivery": prom.ExpectedDelivery,
		"DeliveredOn":      prom.DeliveredOn,
		"Status":           prom.Status,
		"AddedOn":          prom.AddedOn,
	}
	return
}

type DependentPromise struct {
	Parent string
	Child  string
}

func (edge DependentPromise) SourceVertex() string {
	return "Promise"
}

func (edge DependentPromise) TargetVertex() string {
	return "Promise"
}

func (edge DependentPromise) EdgeType() string {
	return "dependent_on"
}

func (edge DependentPromise) TgData() (source string, target string, data tigergraph.TigerGraphObject) {
	source = edge.Child
	target = edge.Parent
	data = make(tigergraph.TigerGraphObject)
	return
}

var _ tigergraph.TigerGraphEdge = new(DependentPromise)
var _ db.MultiDatabaseEntity = new(Promise)
