package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func PostTopSecretMessage(ctx *gin.Context) {
	var request entity.SatellitesRequest

	ctx.ShouldBindJSON(&request)
	errValidation := validate.Struct(request)

	if errValidation == nil {

		answer, err := service.CalculateStarshipCode(request.Satellites)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		} else {
			json.Marshal(answer)
			ctx.JSON(http.StatusOK, answer)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errValidation.Error(),
		})
	}

}

func PostSplittedTopSecretMessage(ctx *gin.Context) {
	name := ctx.Param("name")

	var classifiedMessage entity.Satellite
	classifiedMessage.Name = name
	ctx.ShouldBindJSON(&classifiedMessage)

	if errValidation := validate.Struct(classifiedMessage); errValidation != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errValidation.Error(),
		})
	} else {
		service.PostSatelliteCode(classifiedMessage)
		json.Marshal(classifiedMessage)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "The satellite has been successfully added",
		})
	}
}

func GetSplittedTopSecretMessage(ctx *gin.Context) {

	answer, err := service.GetSplittedSatelliteCode()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else {
		json.Marshal(answer)
		ctx.JSON(http.StatusOK, answer)
	}

}
