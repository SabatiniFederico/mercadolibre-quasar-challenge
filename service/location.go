package service

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/lib/trilateration"
)

func getPositions() []trilateration.Point {

	kenobi := trilateration.Point{X: -500.0, Y: -200.0}
	skywalker := trilateration.Point{X: 100.0, Y: -100.0}
	sato := trilateration.Point{X: 500.0, Y: 100.0}

	return []trilateration.Point{kenobi, skywalker, sato}
}

func GetLocation(distances ...float64) (x, y float64, err error) {

	pos := getPositions()

	solution, err := trilateration.Solve2DTrilateration(pos, distances)

	return solution.X, solution.Y, err
}
