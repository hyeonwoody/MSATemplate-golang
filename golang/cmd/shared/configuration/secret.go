package configuration

import (
	hub "crave/hub/cmd/configuration"
)

type Database struct {
	Uri      string
	Username string
	Password string
}

type Variable struct {
	Database    Database
	HubVariable hub.Variable
}

func NewVariable() *Variable {

	database := &Database{
		Uri:      "bolt://localhost:7687",
		Username: "neo4j",
		Password: "password",
	}

	hubVariable := hub.NewVariable()

	return &Variable{
		Database:    *database,
		HubVariable: *hubVariable,
	}
}
