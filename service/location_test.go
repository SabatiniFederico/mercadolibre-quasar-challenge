package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSatellites(t *testing.T) {
	if sats[0].Name != "Kenobi" {
		t.Error("Name kenobi is not present")
	}
}

func TestGetPosition(t *testing.T) {
	var tests = []struct {
		inputDistances []float32
		expectedX      float32
		expectedY      float32
	}{
		{
			inputDistances: []float32{650, 350, 680},
			expectedX:      -13.513513565063477,
			expectedY:      231.08108520507812,
		},
		{
			inputDistances: []float32{650, 350, 530},
			expectedX:      -13.513513565063477,
			expectedY:      231.08108520507812,
		},
		{
			inputDistances: []float32{643.3333, 450, 633.3333},
			expectedX:      -100.0,
			expectedY:      300.0,
		},
		{
			inputDistances: []float32{725, 300, 406.6666},
			expectedX:      100.0,
			expectedY:      200.0,
		},
		{
			inputDistances: []float32{924, 600, 566.6666},
			expectedX:      100.0,
			expectedY:      500.0,
		},
		{
			inputDistances: []float32{300, 632.5, 1000},
			expectedX:      -500.0,
			expectedY:      100.0,
		},
	}

	for _, test := range tests {
		x, y, err := GetLocation(test.inputDistances...)

		assert.Nil(t, err)
		assert.InDelta(t, test.expectedX, x, 5)
		assert.InDelta(t, test.expectedY, y, 5)
	}
}