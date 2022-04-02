package trilateration

import (
	"math"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

func Solve2DTrilateration(p1 model.Point, p2 model.Point, p3 model.Point, distances []float64) model.Point {

	coordinateX := (math.Pow(distances[0], 2) - math.Pow(distances[1], 2) + math.Pow(p2.X, 2)) / (2 * p2.X)
	coordinateY := math.Sqrt(math.Pow(distances[0], 2) - math.Pow(coordinateX, 2))

	return model.Point{X: coordinateX, Y: coordinateY}
}
