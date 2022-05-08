package location

import (
	"github.com/yourwordquest/core/common"
	"github.com/yourwordquest/core/db"
	"github.com/yourwordquest/core/db/es"
	"github.com/yourwordquest/core/db/graph"
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

func (loc *Location) GraphUpdate() (entities []graph.GraphEntity) {
	node := graph.GraphEntity{
		Node: &graph.Node{
			Id:   loc.id,
			Name: "Location",
		},
		Data: map[string]interface{}{
			"Name":           loc.Name,
			"Code":           loc.Code,
			"GovernmentType": loc.GovernmentType,
			"Status":         loc.Status,
			"Classification": loc.Classification,
		},
	}
	entities = []graph.GraphEntity{node}

	for i := range loc.Parents {
		child_edge := graph.GraphEntity{
			Edge: &graph.Edge{
				Name:          "child_location_of",
				Source:        "Location",
				SourceId:      loc.Parents[i],
				Destination:   "Location",
				DestinationId: loc.id,
			},
		}
		parent_edge := graph.GraphEntity{
			Edge: &graph.Edge{
				Name:          "parent_location_of",
				Source:        "Location",
				SourceId:      loc.id,
				Destination:   "Location",
				DestinationId: loc.Parents[i],
			},
		}

		entities = append(entities, child_edge, parent_edge)
	}
	return
}

var _ db.FirestoreDocument = new(Location)
var _ es.ElasticSearchDocument = new(Location)
var _ graph.Neo4JDocument = new(Location)
