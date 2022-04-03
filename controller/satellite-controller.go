package controller

import (
	"encoding/json"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/service"
	"github.com/gin-gonic/gin"
)

func TopSecretMessage(ctx *gin.Context) {
	var info entity.SatellitesRequest

	err := ctx.ShouldBindJSON(&info)
	if err == nil {
		satellites := info.Satellites

		answer, err := service.GetStarshipClassifiedCode(satellites)

		if err != nil {
			ctx.JSON(404, gin.H{
				"message": "impossible to determine classified message",
			})
		} else {
			json.Marshal(answer)
			ctx.JSON(200, answer)
		}
	} else {
		ctx.JSON(400, gin.H{
			"message": "bad format",
		})
	}

}

func SplittedTopSecretMessage(ctx *gin.Context) {
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
