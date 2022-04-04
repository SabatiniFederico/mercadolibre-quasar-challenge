package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostTopSecretMessage(ctx *gin.Context) {
	var request entity.SatellitesRequest

	err := ctx.ShouldBindJSON(&request)
	if err == nil {

		answer, err := service.CalculateStarshipClassifiedCode(request.Satellites)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "impossible to determine classified message",
			})
		} else {
			json.Marshal(answer)
			ctx.JSON(http.StatusOK, answer)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad format",
		})
	}

}

func PostSplittedTopSecretMessage(ctx *gin.Context) {
	name := ctx.Param("name")

	var classifiedMessage entity.Satellite
	classifiedMessage.Name = name
	ctx.ShouldBindJSON(&classifiedMessage)

	var validate = validator.New()

	if errs := validate.Struct(classifiedMessage); errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad format",
		})
	} else {
		service.AddClassifiedCode(classifiedMessage)
		json.Marshal(classifiedMessage)
		ctx.JSON(http.StatusCreated, classifiedMessage)
	}
}

func GetSplittedTopSecretMessage(ctx *gin.Context) {

	answer, err := service.GetSplittedClassifiedCode()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "impossible to determine classified message",
		})
	} else {
		json.Marshal(answer)
		ctx.JSON(http.StatusOK, answer)
	}

}
