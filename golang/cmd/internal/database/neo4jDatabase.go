package database

import (
	"crave/configuration"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var DB neo4j.Driver

func ConnectDatabase(database *configuration.Database) {

	var err error
	DB, err = neo4j.NewDriver(database.Uri, neo4j.BasicAuth(database.Username, database.Password, ""))
	if err != nil {
		log.Fatalf("failed to connect to Neo4j: %v", err)
	}
}
