package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func PostTopSecretMessage(ctx *gin.Context) {
	var request entity.SatellitesRequest

	errParsingJSON := ctx.ShouldBindJSON(&request)
	errValidation := validate.Struct(request)

	fmt.Println(errParsingJSON)
	fmt.Println(errValidation)
	fmt.Println(request)

	if errParsingJSON == nil && errValidation == nil {

		answer, err := service.CalculateStarshipCode(request.Satellites)

		if err == nil {
			json.Marshal(answer)
			ctx.JSON(http.StatusOK, answer)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error ocurred when parsing satellites json",
		})
	}

}

func PostSplittedTopSecretMessage(ctx *gin.Context) {
	name := ctx.Param("name")

	var classifiedMessage entity.Satellite
	classifiedMessage.Name = name

	errParsingJSON := ctx.ShouldBindJSON(&classifiedMessage)
	errValidation := validate.Struct(classifiedMessage)

	if errParsingJSON == nil && errValidation == nil {
		service.AddSatelliteCode(classifiedMessage)
		json.Marshal(classifiedMessage)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "The satellite has been successfully added",
		})
	} else {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error ocurred when parsing satellites json",
		})
	}
}

func GetSplittedTopSecretMessage(ctx *gin.Context) {

	answer, err := service.GetStarshipCodeFromStoredSatellites()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else {
		json.Marshal(answer)
		ctx.JSON(http.StatusOK, answer)
	}

}
