package main

import (
	"crave/hub/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	go startLib(router)
	router.Run(":3001")
}

func startLib(router *gin.Engine) {
	go lib.Start(router)
}
