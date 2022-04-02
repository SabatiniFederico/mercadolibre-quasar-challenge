package trilateration

import (
	"fmt"
	"math"
	"testing"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

func TestPreciseTrilateration(t *testing.T) {
	pos1 := model.Point{X: 0, Y: 0}
	pos2 := model.Point{X: 100, Y: 0}
	pos3 := model.Point{X: 50, Y: 30}
	dists := []float64{50, 50, 30}

	actualPosition, error := Solve2DTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := model.Point{X: 50.0, Y: 0.0}

	if error != nil {
		t.Errorf("There was no solution for expected position (%f, %f)", actualPosition.X, actualPosition.Y)
	}
	fmt.Printf("(%f %f)", actualPosition.X, actualPosition.Y)
	if actualPosition != expectedPosition {
		t.Errorf("The position expected was (%f, %f), but received (%f, %f)", actualPosition.X, actualPosition.Y, expectedPosition.X, expectedPosition.Y)
	}
}

func TestNonPreciseTrilateration(t *testing.T) {
	pos1 := model.Point{X: 0, Y: 0}
	pos2 := model.Point{X: 100, Y: 0}
	pos3 := model.Point{X: 50, Y: 30}
	dists := []float64{100, 100, 55}

	actualPosition, error := Solve2DTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := model.Point{X: 50.0, Y: 86.6025}

	if error != nil {
		t.Errorf("There was no solution for expected position (%f, %f)", actualPosition.X, actualPosition.Y)
	}

	if isNonAccuratePosition(actualPosition, expectedPosition) {
		t.Errorf("The position expected was (%f, %f), but received (%f, %f)", actualPosition.X, actualPosition.Y, expectedPosition.X, expectedPosition.Y)
	}
}

func TestNoSolutionForTrilateration(t *testing.T) {
	pos1 := model.Point{X: 0, Y: 0}
	pos2 := model.Point{X: 100, Y: 0}
	pos3 := model.Point{X: 50, Y: 30}
	dists := []float64{10, 10, 10}

	actualPosition, error := Solve2DTrilateration(pos1, pos2, pos3, dists)

	if error == nil {
		t.Errorf("There was a solution for imposible trilateration case (%f, %f)", actualPosition.X, actualPosition.Y)
	}
}

func isNonAccuratePosition(actualPoint model.Point, expectedPoint model.Point) bool {
	if math.Abs(actualPoint.X-expectedPoint.X) > marginOfError {
		return true
	}
	if math.Abs(actualPoint.Y-expectedPoint.Y) > marginOfError {
		return true
	}
	return false
}
