package main

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/topsecret/", controller.TopSecretMessage)
	server.POST("/topsecret_split/:name", controller.TopSecretSplittedMessage)

	server.Run(":8080")
}
