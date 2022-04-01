package trilateration

import (
	"math"
)

type position struct {
	X int32
	Y int32
}

type satellite struct {
	name string
	pos  position
}

var sattellites = []satellite{
	satellite{
		name: "Kenobi",
		pos:  position{-500, -200},
	},
	satellite{
		name: "Skywalker",
		pos:  position{100, -100},
	},
	satellite{
		name: "Sato",
		pos:  position{500, 100},
	},
}

func GetLocation(distances ...float32) (x, y float32) {
	//var coordinateX float32
	//var coordinateY float32

	var squaredDistances []float32

	for _, distance := range distances {
		squaredDistances = append(squaredDistances, float32(math.Pow(float64(distance), 2)))
	}

	return 20, 20
}

func GetMessage(messages ...[]string) (msg string) {
	return "some message"
}
