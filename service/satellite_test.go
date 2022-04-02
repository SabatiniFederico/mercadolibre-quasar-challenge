package service

import (
	"testing"
)

func TestSatellites(t *testing.T) {
	if sats[0].Name != "Kenobi" {
		t.Error("Name kenobi is not present")
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
