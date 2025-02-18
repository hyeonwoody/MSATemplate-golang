package main

import (
	configuration "crave/configuration"
	hubConfiguration "crave/internal/configuration/hub"
	"crave/internal/database"
	"crave/internal/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	variable := configuration.NewVariable()
	database.ConnectDatabase(&variable.Database)
	router := gin.Default()
	container := hubConfiguration.NewContainer(&variable.HubVariable, &database.DB, router)

	go startApiLib(container)
	router.Run(":3000")
}

func startApiLib(container *hubConfiguration.Container) {
	go lib.StartHubLib(container)
}
