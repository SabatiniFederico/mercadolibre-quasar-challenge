package main

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.POST("/topsecret/", controller.PostTopSecretMessage)
	server.POST("/topsecret_split/:name", controller.PostSplittedTopSecretMessage)
	server.GET("/topsecret_split/", controller.GetSplittedTopSecretMessage)

	server.Run(":8080")
}
