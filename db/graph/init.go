package graph

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/yourwordquest/core/utils"
	"github.com/yourwordquest/core/utils/secrets"
)

var session neo4j.Session
var driver neo4j.Driver

func init() {
	username, err := secrets.GetSecret("NEO4J_USERNAME")
	utils.MaybePanic(err)
	password, err := secrets.GetSecret("NEO4J_PASSWORD")
	utils.MaybePanic(err)
	uri, err := secrets.GetSecret("NEO4J_URI")
	utils.MaybePanic(err)

	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	utils.MaybePanic(err)
	utils.MaybePanic(driver.VerifyConnectivity())

	session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func CloseDB() {
	driver.Close()
	session.Close()
}
