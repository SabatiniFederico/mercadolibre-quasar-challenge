package main

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.POST("/topsecret/", controller.TopSecretMessage)
	server.POST("/topsecret_split/:satellite_name", controller.SplittedTopSecretMessage)
	server.GET("/topsecret_split/", nil)

	server.Run(":8080")
}
