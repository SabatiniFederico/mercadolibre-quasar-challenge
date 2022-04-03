package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessageWithNoPhaseShift(t *testing.T) {

	message1 := []string{"este", "es", "un", ""}
	message2 := []string{"", "", "un", "mensaje"}
	message3 := []string{"", "es", "", ""}

	actualString, err := GetMessage(message1, message2, message3)
	expectedString := "este es un mensaje"

	assert.Nil(t, err)
	if actualString != expectedString {
		t.Errorf("actual string is: %s but expected message was: %s", actualString, expectedString)
	}
}
