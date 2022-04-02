package service

import (
	"testing"
)

func TestSatellites(t *testing.T) {
	if sats[0].Name != "Kenobi" {
		t.Error("Name kenobi is not present")
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
