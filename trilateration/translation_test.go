package trilateration

import (
	"math"
	"testing"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

func TestTranslatePoints(t *testing.T) {

	inputPoints := []model.Point{{X: -500, Y: -200}, {X: 100, Y: -100}, {X: 500, Y: 100}}
	translation := model.Point{X: 500, Y: 200}
	expected := []model.Point{{X: 0, Y: 0}, {X: 600, Y: 100}, {X: 1000, Y: 300}}

	actualPoints := translatePoints(translation, inputPoints...)

	for i, expectedPoint := range expected {
		if isNonAccuratePosition(expectedPoint, actualPoints[i]) {
			t.Errorf("The position expected was (%f, %f), but received (%f, %f)", expectedPoint.X, expectedPoint.Y, actualPoints[i].X, actualPoints[i].Y)
		}
	}
}

func TestRotatePoints(t *testing.T) {

	var tests = []struct {
		inputPoints   []model.Point
		inputRotation float64
		expected      []model.Point
	}{
		{
			inputPoints:   []model.Point{{X: 0, Y: 0}, {X: 600, Y: 100}, {X: 1000, Y: 300}},
			inputRotation: math.Pi / 2,
			expected:      []model.Point{{X: 0, Y: 0}, {X: -100, Y: 600}, {X: -300, Y: 1000}},
		},
		{
			inputPoints:   []model.Point{{X: 0, Y: 0}, {X: 600, Y: 100}, {X: 1000, Y: 300}},
			inputRotation: math.Pi,
			expected:      []model.Point{{X: 0, Y: 0}, {X: -600, Y: -100}, {X: -1000, Y: -300}},
		},
	}

	for _, test := range tests {
		actualPoints := rotatePoints(test.inputRotation, test.inputPoints...)
		for i, expectedPoint := range test.expected {
			if isNonAccuratePosition(expectedPoint, actualPoints[i]) {
				t.Errorf("The position expected was (%f, %f), but received (%f, %f)", expectedPoint.X, expectedPoint.Y, actualPoints[i].X, actualPoints[i].Y)
			}
		}
	}

}
