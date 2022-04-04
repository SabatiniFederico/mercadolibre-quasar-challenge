package service

import (
	"errors"
	"fmt"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

var storedSatellites []entity.Satellite

func AddClassifiedCode(newClassifiedMessage entity.Satellite) {
	for i, satellite := range storedSatellites {
		if satellite.Name == newClassifiedMessage.Name {
			storedSatellites[i] = newClassifiedMessage
			return
		}
	}
	storedSatellites = append(storedSatellites, newClassifiedMessage)
}

func GetSplittedClassifiedCode() (entity.StarshipResponse, error) {
	return CalculateStarshipClassifiedCode(storedSatellites)
}

func CalculateStarshipClassifiedCode(satellites []entity.Satellite) (answer entity.StarshipResponse, err error) {

	distances, criptedMessages, errRequest := getValuesFromSatellites(satellites)

	x, y, errLocation := GetLocation(distances[0], distances[1], distances[2])
	message, errMessage := GetMessage(criptedMessages[0], criptedMessages[1], criptedMessages[2])

	if errLocation != nil || errMessage != nil || errRequest != nil {
		fmt.Print("could not calculate position or message with provided info \n")
		return entity.StarshipResponse{}, errors.New("could not calculate position or message with provided info")
	}

	answer = entity.StarshipResponse{
		Position: entity.Position{X: float64(x), Y: float64(y)},
		Message:  message,
	}
	return answer, nil
}

func getValuesFromSatellites(satellites []entity.Satellite) (distances [3]float64, messages [3][]string, err error) {

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

func hasRepeatedSatelliteNames(satellites []entity.Satellite) bool {
	for i, satellite := range satellites {
		for j := i + 1; j < len(satellites); j++ {
			if satellite.Name == satellites[j].Name {
				return true
			}
		}
	}
	return false
}
