package location

import "github.com/microcosm-cc/bluemonday"

var policy *bluemonday.Policy

func init() {
	policy = bluemonday.UGCPolicy()
}
