package service

import (
	"errors"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/data"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

var storedSatellites []entity.Satellite

func AddSatelliteCode(newClassifiedMessage entity.Satellite) {
	for i, satellite := range storedSatellites {
		if satellite.Name == newClassifiedMessage.Name {
			storedSatellites[i] = newClassifiedMessage
			return
		}
	}
	storedSatellites = append(storedSatellites, newClassifiedMessage)
}

func GetStarshipCodeFromStoredSatellites() (entity.StarshipResponse, error) {
	return CalculateStarshipCode(storedSatellites)
}

func CalculateStarshipCode(satellites []entity.Satellite) (answer entity.StarshipResponse, err error) {

	distances, criptedMessages, errRequest := getValuesFromSatellites(satellites)

	if errRequest != nil {
		return answer, errRequest
	}

	x, y, errLocation := GetLocation(distances...)
	message, errMessage := GetMessage(criptedMessages...)

	if errLocation != nil || errMessage != nil {
		return answer, errors.New("couldn't calculate position or message with provided info")
	}

	answer = entity.StarshipResponse{
		Position: entity.Position{X: float64(x), Y: float64(y)},
		Message:  message,
	}
	return answer, nil
}

func getValuesFromSatellites(satellites []entity.Satellite) (distances []float64, messages [][]string, err error) {

	if len(satellites) != 3 || hasRepeatedSatelliteNames(satellites) {
		return distances, messages, errors.New("request is invalid, we expected to receive kenobi, sato, and skywalker once")
	}

	sortedSatellites := sortSatellites(satellites)

	for _, satellite := range sortedSatellites {
		messages = append(messages, satellite.Message)
		distances = append(distances, satellite.Distance)
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

func sortSatellites(satellites []entity.Satellite) (sortedSats []entity.Satellite) {

	dicc := make(map[string]entity.Satellite)
	for _, satellite := range satellites {
		dicc[satellite.Name] = satellite
	}

	for _, onlineSatellite := range data.GetOnlineSatellites() {
		sortedSats = append(sortedSats, dicc[onlineSatellite.Name])
	}

	return sortedSats
}
