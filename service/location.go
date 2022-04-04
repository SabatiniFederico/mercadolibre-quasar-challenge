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

func GetLocation(distances ...float32) (x, y float32, err error) {

	pos := getPositions()

	distanceKenobi := float64(distances[0])
	distanceSkywalker := float64(distances[1])
	distanceSato := float64(distances[2])

	preciseDistances := []float64{distanceKenobi, distanceSkywalker, distanceSato}

	solution, err := trilateration.Solve2DTrilateration(pos["kenobi"], pos["skywalker"], pos["sato"], preciseDistances)

	return float32(solution.X), float32(solution.Y), err
}
