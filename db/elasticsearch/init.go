package elasticsearch

import "os"

var server = os.Getenv("ES_SERVER")
var username = os.Getenv("ES_USERNAME")
var password = os.Getenv("ES_PASSWORD")
