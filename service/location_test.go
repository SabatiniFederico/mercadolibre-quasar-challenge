package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosition(t *testing.T) {
	var tests = []struct {
		inputDistances []float64
		expectedX      float64
		expectedY      float64
	}{
		{
			inputDistances: []float64{650, 350, 680},
			expectedX:      100.0,
			expectedY:      -450.0,
		},
		{
			inputDistances: []float64{650, 350, 530},
			expectedX:      -13.50,
			expectedY:      231.00,
		},
		{
			inputDistances: []float64{643.3333, 450, 633},
			expectedX:      -100.0,
			expectedY:      303.0,
		},
		{
			inputDistances: []float64{722, 300, 411.2},
			expectedX:      100.0,
			expectedY:      200.0,
		},
		{
			inputDistances: []float64{922, 600, 565.6},
			expectedX:      100.0,
			expectedY:      500.0,
		},
		{
			inputDistances: []float64{300, 632.5, 1000},
			expectedX:      -500.0,
			expectedY:      100.0,
		},
	}

	for _, test := range tests {
		x, y, err := GetLocation(test.inputDistances...)

		assert.Nil(t, err)
		assert.InDelta(t, test.expectedX, x, 2)
		assert.InDelta(t, test.expectedY, y, 2)
	}
}

func TestGetPositionWhenNoSolution(t *testing.T) {
	inputDistances := []float64{100, 100, 100}

	_, _, err := GetLocation(inputDistances...)

	assert.NotNil(t, err)
}
