package service

import (
	"errors"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

func GetStarshipClassifiedCode(satellites [3]entity.ClassifiedMessage) (answer entity.SolutionResponse, err error) {

	x, y, errLocation := GetLocation(satellites[0].Distance, satellites[1].Distance, satellites[2].Distance)
	message, errMessage := GetMessage(satellites[0].Message, satellites[1].Message, satellites[2].Message)

	if errLocation != nil || errMessage != nil {
		return entity.SolutionResponse{}, errors.New("could not calculate position or message with provided info")
	}

	answer = entity.SolutionResponse{Position: entity.Point{X: float64(x), Y: float64(y)}, Message: message}
	return answer, nil
}
