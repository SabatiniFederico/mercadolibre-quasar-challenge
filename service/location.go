package service

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/lib/trilateration"
)

func getPositions() map[string]trilateration.Point {

	positions := make(map[string]trilateration.Point)
	positions["kenobi"] = trilateration.Point{X: -500.0, Y: -200.0}
	positions["skywalker"] = trilateration.Point{X: 100.0, Y: -100.0}
	positions["sato"] = trilateration.Point{X: 500.0, Y: 100.0}

	return positions
}

func GetLocation(distances ...float64) (x, y float64, err error) {

	pos := getPositions()

	solution, err := trilateration.Solve2DTrilateration(pos["kenobi"], pos["skywalker"], pos["sato"], distances)

	return solution.X, solution.Y, err
}
