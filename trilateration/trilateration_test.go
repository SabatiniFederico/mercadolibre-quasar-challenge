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

	actualPosition := Solve2DTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := model.Point{X: 50.0, Y: 0.0}

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

	actualPosition := Solve2DTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := model.Point{X: 50.0, Y: 86.6025}

	if isNonAccuratePosition(actualPosition, expectedPosition) {
		t.Errorf("The position expected was (%f, %f), but received (%f, %f)", actualPosition.X, actualPosition.Y, expectedPosition.X, expectedPosition.Y)
	}
}

func isNonAccuratePosition(actualPoint model.Point, expectedPoint model.Point) bool {
	margin := 0.0001
	if math.Abs(actualPoint.X-expectedPoint.X) > margin {
		return true
	}
	if math.Abs(actualPoint.Y-expectedPoint.Y) > margin {
		return true
	}
	return false
}
