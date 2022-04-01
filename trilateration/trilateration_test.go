package trilateration

import (
	"testing"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

func TestSolveTrilateration(t *testing.T) {
	pos1 := model.Position{0, 0}
	pos2 := model.Position{100, 0}
	pos3 := model.Position{50, 30}
	dists := []float32{50, 50, 30}

	actualPosition := solveTrilateration(pos1, pos2, pos3, dists)
	expectedPosition := model.Position{50.0, 0.0}

	if actualPosition != expectedPosition {
		t.Errorf("The position expected was (%f, %f), but received (%f, %f)", actualPosition.X, actualPosition.Y, expectedPosition.X, expectedPosition.Y)
	}
}
