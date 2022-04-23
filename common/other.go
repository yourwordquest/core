package common

type OtherDataItem struct {
	Payload interface{} `json:"payload" firestore:"payload"`
	Type    string      `json:"type" firestore:"type"`
}

type OtherData map[string]OtherDataItem

func (data OtherData) Get(key string, defaultValue interface{}) interface{} {
	item, isSet := data[key]
	if !isSet {
		return defaultValue
	}
	return item.Payload
}
