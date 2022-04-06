package service

import (
	"testing"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
	"github.com/stretchr/testify/assert"
)

func getTestData() []entity.Satellite {
	return []entity.Satellite{
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
}

func TestCalculateStarshipCode(t *testing.T) {

	testInput := getTestData()
	expectedOutput := "este es un challenge de mercado libre"
	solution, err := CalculateStarshipCode(testInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, solution.Message)

}

func TestCalculateStarshipCodeWithInvalidDistance(t *testing.T) {

	testInput := getTestData()
	testInput[0].Distance = 100

	_, err := CalculateStarshipCode(testInput)

	assert.NotNil(t, err)

}

func TestCalculateStarshipCodeWithModifiedOrder(t *testing.T) {

	testData := getTestData()
	testInput := []entity.Satellite{
		testData[2],
		testData[0],
		testData[1],
	}

	expectedOutput := "este es un challenge de mercado libre"
	solution, err := CalculateStarshipCode(testInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, solution.Message)

}

func TestCalculateStarshipCodeWithRepeatedKenobi(t *testing.T) {

	testData := getTestData()
	testInput := []entity.Satellite{
		testData[0],
		testData[0],
		testData[0],
	}

	_, err := CalculateStarshipCode(testInput)
	assert.NotNil(t, err)
}

func TestCalculateStarshipCodeWithoutEnoughSatellites(t *testing.T) {

	testData := getTestData()
	testInput := []entity.Satellite{testData[0]}

	_, err := CalculateStarshipCode(testInput)
	assert.NotNil(t, err)
}

func TestAddSatelliteCodeTwoTimesWithTheSameSatellite(t *testing.T) {

	testData := getTestData()
	AddSatelliteCode(testData[0])
	AddSatelliteCode(testData[0])

	assert.Equal(t, 1, len(storedSatellites))
}

func TestGetSplittedStarshipCodeWithNoAdditions(t *testing.T) {

	_, err := GetStarshipCodeFromStoredSatellites()
	assert.NotNil(t, err)
}

func TestGetSplittedStarshipCodeWithCorrectAditions(t *testing.T) {

	testData := getTestData()
	expectedOutput := "este es un challenge de mercado libre"

	AddSatelliteCode(testData[0])
	AddSatelliteCode(testData[1])
	AddSatelliteCode(testData[2])

	solution, err := GetStarshipCodeFromStoredSatellites()

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, solution.Message)
}
