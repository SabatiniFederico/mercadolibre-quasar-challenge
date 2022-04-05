package service

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/lib/trilateration"
)

func getPositions() (positions []trilateration.Point) {

	positions = append(positions, trilateration.Point{X: -500.0, Y: -200.0})
	positions = append(positions, trilateration.Point{X: 100.0, Y: -100.0})
	positions = append(positions, trilateration.Point{X: 500.0, Y: 100.0})

	return positions
}

func GetLocation(distances ...float64) (x, y float64, err error) {

	positions := getPositions()

	solution, err := trilateration.Solve2DTrilateration(positions, distances)

	return solution.X, solution.Y, err
}
