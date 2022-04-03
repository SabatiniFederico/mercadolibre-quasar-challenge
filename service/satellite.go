package service

import (
	"errors"
	"strings"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/trilateration"
)

var sats = []entity.Satellite{
	{
		Name: "Kenobi",
		Pos:  entity.Point{X: -500, Y: -200},
	},
	{
		Name: "Skywalker",
		Pos:  entity.Point{X: 100, Y: -100},
	},
	{
		Name: "Sato",
		Pos:  entity.Point{X: 500, Y: 100},
	},
}

func GetLocation(distances ...float32) (x, y float32, err error) {

	preciseDistances := []float64{float64(distances[0]), float64(distances[1]), float64(distances[2])}
	solution, err := trilateration.Solve2DTrilateration(sats[0].Pos, sats[1].Pos, sats[2].Pos, preciseDistances)

	return float32(solution.X), float32(solution.Y), err
}

func GetMessage(messages ...[]string) (msg string, err error) {

	len := len(messages[0])
	solution := []string{}

	for i := 0; i < len; i++ {
		word, err := mergeTwoWords(messages[0][i], messages[1][i])

		if err != nil {
			return "", err
		}

		word, err2 := mergeTwoWords(word, messages[2][i])

		if err2 != nil {
			return "", err2
		}

		solution = append(solution, word)
	}

	return strings.Join(solution, " "), nil
}
func mergeTwoWords(word1, word2 string) (string, error) {
	if word1 != "" {
		if word1 != word2 && word2 != "" {
			return word1, errors.New("contradiction found")
		}
		return word1, nil
	}
	return word2, nil
}
