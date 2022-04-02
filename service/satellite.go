package service

import (
	"fmt"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/trilateration"
)

var sats = []model.Satellite{
	{
		Name: "Kenobi",
		Pos:  model.Point{X: -500, Y: -200},
	},
	{
		Name: "Skywalker",
		Pos:  model.Point{X: 100, Y: -100},
	},
	{
		Name: "Sato",
		Pos:  model.Point{X: 500, Y: 100},
	},
}

func GetLocation(distances ...float32) (x, y float32) {

	preciseDistances := []float64{float64(distances[0]), float64(distances[1]), float64(distances[2])}
	solution, err := trilateration.Solve2DTrilateration(sats[0].Pos, sats[1].Pos, sats[2].Pos, preciseDistances)

	if err != nil {
		//TODO replace with logger
		fmt.Printf("no hay solucion")
	}
	return float32(solution.X), float32(solution.Y)
}

func GetMessage(messages ...[]string) (msg string) {
	return "some message"
}
