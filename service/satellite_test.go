package service

import (
	"testing"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/stretchr/testify/assert"
)

var testSatellites = []entity.Satellite{
	{
		Name:     "kenobi",
		Distance: 650.0,
		Message:  []string{"", "es", "", "challenge", "", "mercado", "libre"},
	},
	{
		Name:     "skywalker",
		Distance: 350.0,
		Message:  []string{"este", "es", "", "", "de", "", "libre"},
	},
	{
		Name:     "sato",
		Distance: 680.0735,
		Message:  []string{"", "", "", "este", "", "un", "", "de", "mercado", "libre"},
	},
}

func TestCalculateStarshipCode(t *testing.T) {

	expectedOutput := "este es un challenge de mercado libre"
	solution, err := CalculateStarshipCode(testSatellites)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, solution.Message)

}

func TestCalculateStarshipCodeWithModifiedOrder(t *testing.T) {

	testInput := []entity.Satellite{
		testSatellites[2],
		testSatellites[0],
		testSatellites[1],
	}

	expectedOutput := "este es un challenge de mercado libre"
	solution, err := CalculateStarshipCode(testInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, solution.Message)

}

func TestCalculateStarshipCodeWithRepeatedKenobi(t *testing.T) {

	testInput := []entity.Satellite{
		testSatellites[0],
		testSatellites[0],
		testSatellites[0],
	}

	_, err := CalculateStarshipCode(testInput)
	assert.NotNil(t, err)
}

func TestCalculatedStarshipWithNoEnoughSatellites(t *testing.T) {

	testInput := []entity.Satellite{testSatellites[0]}

	_, err := CalculateStarshipCode(testInput)
	assert.NotNil(t, err)
}

func TestAddSatelliteCodeTwoTimesWithTheSameSatellite(t *testing.T) {

	AddSatelliteCode(testSatellites[0])
	AddSatelliteCode(testSatellites[0])

	assert.Equal(t, 1, len(storedSatellites))
}

func TestGetSplittedSatelliteCodeWithNoAdditions(t *testing.T) {

	_, err := GetStarshipCodeFromStoredSatellites()
	assert.NotNil(t, err)
}

func TestGetSplittedClassifiedCodeWithCorrectAditions(t *testing.T) {

	expectedOutput := "este es un challenge de mercado libre"

	AddSatelliteCode(testSatellites[0])
	AddSatelliteCode(testSatellites[1])
	AddSatelliteCode(testSatellites[2])

	solution, err := GetStarshipCodeFromStoredSatellites()

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, solution.Message)
}
