package service

import (
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

	distanceKenobi := float64(distances[0])
	distanceSkywalker := float64(distances[1])
	distanceSato := float64(distances[2])

	preciseDistances := []float64{distanceKenobi, distanceSkywalker, distanceSato}
	solution, err := trilateration.Solve2DTrilateration(sats[0].Pos, sats[1].Pos, sats[2].Pos, preciseDistances)

	return float32(solution.X), float32(solution.Y), err
}
