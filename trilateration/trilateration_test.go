package trilateration

import (
	"math"
	"testing"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

func TestPreciseTrilateration(t *testing.T) {
	pos1 := entity.Point{X: 0, Y: 0}
	pos2 := entity.Point{X: 100, Y: 0}
	pos3 := entity.Point{X: 50, Y: 30}
	dists := []float64{50, 50, 30}

	actualPosition, error := Solve2DTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := entity.Point{X: 50.0, Y: 0.0}

	if error != nil {
		t.Errorf("There was no solution for expected position (%f, %f)", actualPosition.X, actualPosition.Y)
	}

	if actualPosition != expectedPosition {
		t.Errorf("The position expected was (%f, %f), but received (%f, %f)", actualPosition.X, actualPosition.Y, expectedPosition.X, expectedPosition.Y)
	}
}

func TestNonPreciseTrilateration(t *testing.T) {
	pos1 := entity.Point{X: 0, Y: 0}
	pos2 := entity.Point{X: 100, Y: 0}
	pos3 := entity.Point{X: 50, Y: 30}
	dists := []float64{100, 100, 56.6025}

	actualPosition, error := Solve2DTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := entity.Point{X: 50.0, Y: 86.6025}

	if error != nil {
		t.Errorf("There was no solution for expected position (%f, %f)", actualPosition.X, actualPosition.Y)
	}

	if isNonAccuratePosition(actualPosition, expectedPosition) {
		t.Errorf("The position expected was (%f, %f), but received (%f, %f)", expectedPosition.X, expectedPosition.Y, actualPosition.X, actualPosition.Y)
	}
}

func TestNoSolutionForTrilateration(t *testing.T) {
	pos1 := entity.Point{X: 0, Y: 0}
	pos2 := entity.Point{X: 100, Y: 0}
	pos3 := entity.Point{X: 50, Y: 30}
	dists := []float64{10, 10, 10}

	actualPosition, error := Solve2DTrilateration(pos1, pos2, pos3, dists)

	if error == nil {
		t.Errorf("There was a solution for imposible trilateration case (%f, %f)", actualPosition.X, actualPosition.Y)
	}
}

func isNonAccuratePosition(actualPoint entity.Point, expectedPoint entity.Point) bool {
	if math.Abs(actualPoint.X-expectedPoint.X) > 0.0001 {
		return true
	}
	if math.Abs(actualPoint.Y-expectedPoint.Y) > 0.0001 {
		return true
	}
	return false
}
