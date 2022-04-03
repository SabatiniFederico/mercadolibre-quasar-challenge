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
			inputDistances: []float32{643.3333, 450, 633.3333},
			expectedX:      -100.0,
			expectedY:      300.0,
		},
		{
			inputDistances: []float32{716.6666, 300, 406.6666},
			expectedX:      100.0,
			expectedY:      200.0,
		},
		{
			inputDistances: []float32{916.6666, 600, 566.6666},
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
		x, y := GetLocation(test.inputDistances...)

		//fmt.Printf("case %d (x: %f, y: %f), expected: (x: %f, y: %f) \n", i, x, y, test.expectedX, test.expectedY)
		assert.InDelta(t, test.expectedX, x, 10)
		assert.InDelta(t, test.expectedY, y, 10)
	}
}
func TestGetMessageWithNoPhaseShift(t *testing.T) {

	message1 := []string{"este", "es", "un", ""}
	message2 := []string{"", "", "un", "mensaje"}
	message3 := []string{"", "es", "", ""}

	actualString := GetMessage(message1, message2, message3)
	expectedString := "este es un mensaje"

	if actualString != expectedString {
		t.Errorf("actual string is: %s but expected message was: %s", actualString, expectedString)
	}
}
