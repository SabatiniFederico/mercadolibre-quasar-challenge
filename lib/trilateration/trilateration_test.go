package trilateration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreciseTrilateration(t *testing.T) {
	pos := []Point{{X: 0, Y: 0}, {X: 100, Y: 0}, {X: 50, Y: 30}}
	dists := []float64{50, 50, 30}

	actualPosition, err := Solve2DTrilateration(pos, dists)
	expectedPosition := Point{X: 50.0, Y: 0.0}

	assert.Nil(t, err)
	assert.Equal(t, expectedPosition, actualPosition)
}

func TestNonPreciseTrilateration(t *testing.T) {
	pos := []Point{{X: 0, Y: 0}, {X: 100, Y: 0}, {X: 50, Y: 30}}
	dists := []float64{100, 100, 56.6025}

	actualPosition, err := Solve2DTrilateration(pos, dists)
	expectedPosition := Point{X: 50.0, Y: 86.6025}

	assert.Nil(t, err)
	assert.InDelta(t, actualPosition.X, expectedPosition.X, 0.005)
	assert.InDelta(t, actualPosition.Y, expectedPosition.Y, 0.005)
}

func TestNoSolutionForTrilateration(t *testing.T) {
	pos := []Point{{X: 0, Y: 0}, {X: 100, Y: 0}, {X: 50, Y: 30}}
	dists := []float64{10, 10, 10}

	_, err := Solve2DTrilateration(pos, dists)

	assert.NotNil(t, err)

}
