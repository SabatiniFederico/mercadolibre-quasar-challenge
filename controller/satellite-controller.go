package controller

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/gin-gonic/gin"
)

func TopSecretMessage(ctx *gin.Context) {
	var satellites entity.SatellitesInfo

	err := ctx.ShouldBindJSON(&satellites)
	if err == nil {
		ctx.JSON(200, gin.H{
			"position": "hola",
			"message":  "este es un mensaje secreto ",
		})
	} else {
		ctx.JSON(404, gin.H{
			"message": "se a producido un error",
		})
	}

}

func TopSecretSplittedMessage(ctx *gin.Context) {
	name := ctx.Param("name")

	var classifiedMessage entity.ClassifiedMessage
	classifiedMessage.Name = name
	err := ctx.ShouldBindJSON(&classifiedMessage)
	if err == nil {
		ctx.JSON(200, gin.H{
			"position": "hola",
			"message":  "esta es una prueba",
		})
	} else {
		ctx.JSON(404, gin.H{
			"message": "se a producido un error",
		})
	}

}
