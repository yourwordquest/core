package location

import (
	"fmt"

	"github.com/yourwordquest/core/common"
	"github.com/yourwordquest/core/db"
	"github.com/yourwordquest/core/db/elasticsearch"
	"github.com/yourwordquest/core/utils"
)

type Location struct {
	id             string
	Name           string           `firestore:"Name"`
	Code           string           `firestore:"Code"`
	GovernmentType string           `firestore:"GovType"`
	Intro          string           `firestore:"Intro"`
	Narrative      string           `firestore:"Narrative"`
	Other          common.OtherData `firestore:"Other"`
	Status         string           `firestore:"Status"`
	Classification string           `firestore:"Classification"`
	Parents        []string         `firestore:"Parents"`
}

func (loc *Location) Id() string {
	return loc.id
}

func (loc *Location) SetId(id string) {
	loc.id = id
}

func (loc *Location) Collection() string {
	return "Locations"
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

func (loc *Location) GraphData() (queries []string, params map[string]interface{}) {
	params = map[string]interface{}{
		"id":             loc.Id,
		"name":           loc.Name,
		"code":           loc.Code,
		"gov_type":       loc.GovernmentType,
		"status":         loc.Status,
		"classification": loc.Classification,
	}
	queries = []string{
		`MERGE (location:Location {Id: $id})
		 SET location = {Id: $id, Name: $name, Code: $code, GovernmentType: $gov_type, Status: $status, Classification: $classification}
		`,
	}

	for i := range loc.Parents {
		parent_key := fmt.Sprintf("parent_%v", i)
		params[parent_key] = loc.Parents[i]
		query := fmt.Sprintf(`
			MATCH
				(child_location:Location {Id: $id}),
				(parent_location:Location {Id: $%s})
			MERGE (child_location)-[cr:child_location_of]->(parent_location)
			MERGE (parent_location)-[pr:parent_location_of]->(child_location)
		`, parent_key)
		queries = append(queries, query)
	}
	return
}

var _ db.FirestoreDocument = new(Location)
var _ elasticsearch.ElasticSearchDocument = new(Location)
