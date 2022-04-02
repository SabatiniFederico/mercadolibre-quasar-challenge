package controller

import (
	"encoding/json"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/service"
	"github.com/gin-gonic/gin"
)

func TopSecretMessage(ctx *gin.Context) {
	var info entity.SatellitesInfo

	err := ctx.ShouldBindJSON(&info)
	if err == nil {
		satellites := info.Satellites
		x, y := service.GetLocation(satellites[0].Distance, satellites[1].Distance, satellites[2].Distance)
		message := service.GetMessage(satellites[0].Message, satellites[1].Message, satellites[2].Message)

		answer := entity.Solution{Position: entity.Point{X: float64(x), Y: float64(y)}, Message: message}
		json.Marshal(answer)
		ctx.JSON(200, answer)
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
