package utils

import "github.com/microcosm-cc/bluemonday"

var policy *bluemonday.Policy

func init() {
	policy = bluemonday.UGCPolicy()
}

func StripHTML(input string) string {
	return policy.Sanitize(input)
}
