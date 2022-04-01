package triangulation

import (
	"testing"
)

func TestGetLocation(t *testing.T) {

	actualX, actualY := GetLocation(40, 20, 30)
	var expectedX, expectedY float32 = 20, 20

	if expectedX != actualX {
		t.Errorf("Expected Position X (%f) is not same as"+
			" actual string (%f)", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expected Position Y (%f) is not same as"+
			" actual string (%f)", expectedY, actualY)
	}
}

func TestGetMessage(t *testing.T) {

	words := []string{"", "abc", "hola"}

	actualString := GetMessage(words)
	expectedString := "some message"

	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
