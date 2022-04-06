package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessageWithNoNoise(t *testing.T) {

	message1 := []string{"este", "es", "un", ""}
	message2 := []string{"", "", "un", "mensaje"}
	message3 := []string{"", "es", "", ""}

	actualString, err := GetMessage(message1, message2, message3)
	expectedString := "este es un mensaje"

	assert.Nil(t, err)
	assert.Equal(t, expectedString, actualString)
}

func TestGetMessageWithNoise(t *testing.T) {

	message1 := []string{"este", "es", "", "challenge", "de", "", ""}
	message2 := []string{"", "", "", "", "", "", "", "", "", "un", "", "", "", "libre"}
	message3 := []string{"", "es", "", "", "", "mercado", ""}

	actualString, err := GetMessage(message1, message2, message3)
	expectedString := "este es un challenge de mercado libre"

	assert.Nil(t, err)
	assert.Equal(t, expectedString, actualString)
}

func TestRemoveNoiseOfMessages(t *testing.T) {

	message1 := []string{"", "", "", "", "", "", "", "este", "es", "un", ""}
	message2 := []string{"", "", "un", "mensaje"}

	messages := removeNoiseOfMessages(message1, message2)
	assert.Equal(t, len(messages[0]), len(messages[1]))
}
