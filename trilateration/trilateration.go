package trilateration

import (
	"errors"
	"math"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

var ErrNoSolution = errors.New("there is no posible solution")

const marginOfError = 0.0001

func Solve2DTrilateration(p1 model.Point, p2 model.Point, p3 model.Point, distances []float64) (model.Point, error) {

	coordinateX := (math.Pow(distances[0], 2) - math.Pow(distances[1], 2) + math.Pow(p2.X, 2)) / (2 * p2.X)
	coordinateY := math.Sqrt(math.Pow(distances[0], 2) - math.Pow(coordinateX, 2))

	if math.IsNaN(coordinateY) {
		return model.Point{}, ErrNoSolution
	}

	//i have to check two possible solutions for Y
	if isAccurateDistance(distances[2], math.Pow(p3.X-coordinateX, 2)+math.Pow(p3.Y-coordinateY, 2)) {
		return model.Point{X: coordinateX, Y: coordinateY}, nil
	}
	if isAccurateDistance(distances[2], math.Pow(p3.X-coordinateX, 2)+math.Pow(p3.Y+coordinateY, 2)) {
		return model.Point{X: coordinateX, Y: -coordinateY}, nil
	}

	return model.Point{}, ErrNoSolution
}

func isAccurateDistance(expectedDistance float64, actualDistance float64) bool {
	return math.Abs(expectedDistance-actualDistance) >= marginOfError
}
