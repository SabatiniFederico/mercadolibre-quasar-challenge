package service

import (
	"errors"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

func GetStarshipClassifiedCode(request entity.SatellitesRequest) (answer entity.SolutionResponse, err error) {

	distances, criptedMessages, errRequest := getValuesFromSatellites(request)

	x, y, errLocation := GetLocation(distances[0], distances[1], distances[2])
	message, errMessage := GetMessage(criptedMessages[0], criptedMessages[1], criptedMessages[2])

	if errLocation != nil || errMessage != nil || errRequest != nil {
		return entity.SolutionResponse{}, errors.New("could not calculate position or message with provided info")
	}

	answer = entity.SolutionResponse{Position: entity.Point{X: float64(x), Y: float64(y)}, Message: message}
	return answer, nil
}

func getValuesFromSatellites(request entity.SatellitesRequest) ([3]float32, [3][]string, error) {
	satellites := request.Satellites

	var distances [3]float32
	var messages [3][]string

	if len(satellites) != 3 {
		return distances, messages, errors.New("not enough satellites")
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
