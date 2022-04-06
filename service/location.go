package service

import (
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/data"
	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/lib/trilateration"
)

func getPositions() (positions []trilateration.Point) {

	for _, sat := range data.GetOnlineSatellites() {
		positions = append(positions, trilateration.Point{X: sat.Xcoordinate, Y: sat.Ycoordinate})
	}

	return positions
}

func GetLocation(distances ...float64) (x, y float64, err error) {

	positions := getPositions()

	solution, err := trilateration.Solve2DTrilateration(positions, distances)

	return solution.X, solution.Y, err
}
