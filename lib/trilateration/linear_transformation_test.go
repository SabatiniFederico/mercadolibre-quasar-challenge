package trilateration

import (
	"fmt"
	"math"
	"testing"
)

func TestTranslatePoints(t *testing.T) {

	inputPoints := []Point{{X: -500, Y: -200}, {X: 100, Y: -100}, {X: 500, Y: 100}}
	translation := Point{X: 500, Y: 200}
	expected := []Point{{X: 0, Y: 0}, {X: 600, Y: 100}, {X: 1000, Y: 300}}

	actualPoints := linearPointsTranslation(translation, inputPoints...)

	for i, expectedPoint := range expected {
		if isNonAccuratePosition(expectedPoint, actualPoints[i]) {
			t.Errorf("The position expected was (%f, %f), but received (%f, %f)", expectedPoint.X, expectedPoint.Y, actualPoints[i].X, actualPoints[i].Y)
		}
	}
}

func TestRotateOnePoint(t *testing.T) {
	element := Point{X: 600, Y: 100}

	rotation := math.Acos(element.X / math.Sqrt(math.Pow(element.X, 2)+math.Pow(element.Y, 2)))
	result := rotatePoint(-rotation, rotatePoint(rotation, element))
	fmt.Printf("rotation: (%f, %f), expected (600,100)", result.X, result.Y)
	//if(rotatePoint(-rotation, rotatePoint(rotation, element)))
}
func TestRotatePoints(t *testing.T) {

	var tests = []struct {
		inputPoints   []Point
		inputRotation float64
		expected      []Point
	}{
		{
			inputPoints:   []Point{{X: 0, Y: 0}, {X: 600, Y: 100}, {X: 1000, Y: 300}},
			inputRotation: math.Pi / 2,
			expected:      []Point{{X: 0, Y: 0}, {X: -100, Y: 600}, {X: -300, Y: 1000}},
		},
		{
			inputPoints:   []Point{{X: 0, Y: 0}, {X: 600, Y: 100}, {X: 1000, Y: 300}},
			inputRotation: math.Pi,
			expected:      []Point{{X: 0, Y: 0}, {X: -600, Y: -100}, {X: -1000, Y: -300}},
		},
	}

	for _, test := range tests {
		actualPoints := linearPointsRotation(test.inputRotation, test.inputPoints...)
		for i, expectedPoint := range test.expected {
			if isNonAccuratePosition(expectedPoint, actualPoints[i]) {
				t.Errorf("The position expected was (%f, %f), but received (%f, %f)", expectedPoint.X, expectedPoint.Y, actualPoints[i].X, actualPoints[i].Y)
			}
		}
	}

}
