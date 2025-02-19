package main

import (
	hubLib "crave/hub/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	go startApiLib(router)
	router.Run(":3000")
}

func startApiLib(router *gin.Engine) {
	go hubLib.Start(router)
}
