package router

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/topsecret/", controller.PostTopSecretMessage)
	router.POST("/topsecret_split/:name", controller.PostSplittedTopSecretMessage)
	router.GET("/topsecret_split/", controller.GetSplittedTopSecretMessage)

	return router
}
