package main

import (
	"crave/configuration"
	"crave/internal/database"
	"crave/internal/lib"

	"github.com/gin-gonic/gin"
)

func main() {

	variable := configuration.NewVariable()
	database.ConnectDatabase(&variable.Database)
	router := gin.Default()
	container := configuration.NewContainer(variable, &database.DB, router)

	go startApiLib(container)
	router.Run(":3000")
}

func startApiLib(container *configuration.Container) {
	go lib.StartHubLib(&container.HubContainer)
}
