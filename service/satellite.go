package service

import (
	"errors"
	"fmt"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

var splittedSatellites []entity.ClassifiedMessage

func CalculateStarshipClassifiedCode(satellites []entity.ClassifiedMessage) (answer entity.SolutionResponse, err error) {

	distances, criptedMessages, errRequest := getValuesFromSatellites(satellites)

	x, y, errLocation := GetLocation(distances[0], distances[1], distances[2])
	message, errMessage := GetMessage(criptedMessages[0], criptedMessages[1], criptedMessages[2])

	if errLocation != nil || errMessage != nil || errRequest != nil {
		fmt.Print("could not calculate position or message with provided info \n")
		return entity.SolutionResponse{}, errors.New("could not calculate position or message with provided info")
	}

	answer = entity.SolutionResponse{
		Position: entity.Point{X: float64(x), Y: float64(y)},
		Message:  message,
	}
	return answer, nil
}

func PostSplittedClassifiedCode(newClassifiedMessage entity.ClassifiedMessage) {
	for i, satellites := range splittedSatellites {
		if satellites.Name == newClassifiedMessage.Name {
			splittedSatellites[i] = newClassifiedMessage
			return
		}
	}
	splittedSatellites = append(splittedSatellites, newClassifiedMessage)
}

func GetSplittedClassifiedCode() (answer entity.SolutionResponse, err error) {
	return CalculateStarshipClassifiedCode(splittedSatellites)
}

func getValuesFromSatellites(satellites []entity.ClassifiedMessage) ([3]float32, [3][]string, error) {

	var distances [3]float32
	var messages [3][]string

	if len(satellites) != 3 || hasRepeatedSatelliteNames(satellites) {
		fmt.Print("request information is invalid \n")
		return distances, messages, errors.New("request information is invalid")
	}

	for i := 0; i < len(satellites); i += 1 {
		switch satellites[i].Name {
		case "kenobi":
			distances[0] = satellites[i].Distance
			messages[0] = satellites[i].Message
		case "skywalker":
			distances[1] = satellites[i].Distance
			messages[1] = satellites[i].Message
		case "sato":
			distances[2] = satellites[i].Distance
			messages[2] = satellites[i].Message
		}
	}
	return distances, messages, nil
}

func hasRepeatedSatelliteNames(satellites []entity.ClassifiedMessage) bool {
	for i, satellite := range satellites {
		for j := i + 1; j < len(satellites); j++ {
			if satellite.Name == satellites[j].Name {
				return true
			}
		}
	}
	return false
}
