package location

import (
	"github.com/yourwordquest/core/common"
)

type Location struct {
	Id             string           `firestore:"-,as_key"`
	Name           string           `firestore:"name"`
	Code           string           `firestore:"code"`
	GovernmentType string           `firestore:"govType"`
	Intro          string           `firestore:"intro"`
	Narrative      string           `firestore:"narrative"`
	Other          common.OtherData `firestore:"other"`
	Status         string           `firestore:"status"`
	Classification string           `firestore:"classification"`
}

func (loc Location) EsData() (id string, index string, data map[string]interface{}) {
	// Return data that can be indexed to elastic search
	id = loc.Id
	index = "searchable_locations"
	keywords := loc.Other.Get("keywords", "")

	data = map[string]interface{}{
		"id":             loc.Id,
		"name":           loc.Name,
		"code":           loc.Code,
		"narrative":      policy.Sanitize(loc.Narrative),
		"classification": loc.Classification,
		"keywords":       keywords,
	}
	return
}

func (loc Location) TgData() (id string, data map[string]interface{}) {
	id = loc.Id
	data = map[string]interface{}{
		"Name":           loc.Name,
		"Code":           loc.Code,
		"GovernmentType": loc.GovernmentType,
		"Status":         loc.Status,
		"Classification": loc.Classification,
	}
	return
}
