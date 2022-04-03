package controller

import (
	"encoding/json"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TopSecretMessage(ctx *gin.Context) {
	var request entity.SatellitesRequest

	err := ctx.ShouldBindJSON(&request)
	if err == nil {

		answer, err := service.CalculateStarshipClassifiedCode(request.Satellites)

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

func PostSplittedTopSecretMessage(ctx *gin.Context) {
	name := ctx.Param("name")

	var classifiedMessage entity.ClassifiedMessage
	classifiedMessage.Name = name
	ctx.ShouldBindJSON(&classifiedMessage)

	var validate = validator.New()

	if errs := validate.Struct(classifiedMessage); errs != nil {
		ctx.JSON(400, gin.H{
			"message": "bad format",
		})
	} else {
		service.PostSplittedClassifiedCode(classifiedMessage)
		json.Marshal(classifiedMessage)
		ctx.JSON(200, classifiedMessage)
	}
}

func GetSplittedTopSecretMessage(ctx *gin.Context) {

	answer, err := service.GetSplittedClassifiedCode()
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "impossible to determine classified message",
		})
	} else {
		json.Marshal(answer)
		ctx.JSON(200, answer)
	}

}
