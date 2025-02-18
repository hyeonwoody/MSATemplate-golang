package configuration

import (
	hub "crave/hub/cmd/configuration"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Container struct {
	HubContainer hub.Container
}

func NewContainer(variable *Variable, DB *neo4j.Driver, router *gin.Engine) *Container {

	hubContainer := hub.NewContainer(&variable.HubVariable, DB, router)

	return &Container{
		HubContainer: *hubContainer,
	}
}
