package location

import (
	"github.com/yourwordquest/core/common"
	"github.com/yourwordquest/core/db"
	"github.com/yourwordquest/core/db/tigergraph"
	"github.com/yourwordquest/core/utils"
)

type Location struct {
	id             string
	Name           string           `firestore:"name"`
	Code           string           `firestore:"code"`
	GovernmentType string           `firestore:"govType"`
	Intro          string           `firestore:"intro"`
	Narrative      string           `firestore:"narrative"`
	Other          common.OtherData `firestore:"other"`
	Status         string           `firestore:"status"`
	Classification string           `firestore:"classification"`
}

func (loc *Location) Id() string {
	return loc.id
}

func (loc *Location) SetId(id string) {
	loc.id = id
}

func (loc *Location) Collection() string {
	return "locations"
}

func (loc *Location) ESIndex() string {
	return "searchable_locations"
}

func (loc *Location) EsData() (id string, data map[string]interface{}) {
	// Return data that can be indexed to elastic search
	id = loc.id
	keywords := loc.Other.Get("keywords", "")

	data = map[string]interface{}{
		"id":             loc.Id,
		"name":           loc.Name,
		"code":           loc.Code,
		"narrative":      utils.StripHTML(loc.Narrative),
		"classification": loc.Classification,
		"keywords":       keywords,
	}
	return
}

func (loc *Location) TgVertex() string {
	return "Location"
}

func (loc *Location) TgData() (id string, data tigergraph.TigerGraphObject) {
	id = loc.id
	data = tigergraph.TigerGraphObject{
		"Name":           loc.Name,
		"Code":           loc.Code,
		"GovernmentType": loc.GovernmentType,
		"Status":         loc.Status,
		"Classification": loc.Classification,
	}
	return
}

type ChildLocation struct {
	Parent string
	Child  string
}

func (edge ChildLocation) SourceVertex() string {
	return "Location"
}

func (edge ChildLocation) TargetVertex() string {
	return "Location"
}

func (edge ChildLocation) EdgeType() string {
	return "child_location"
}

func (edge ChildLocation) TgData() (source string, target string, data tigergraph.TigerGraphObject) {
	source = edge.Parent
	target = edge.Child
	data = make(tigergraph.TigerGraphObject)
	return
}

var _ tigergraph.TigerGraphEdge = new(ChildLocation)
var _ db.MultiDatabaseEntity = new(Location)
