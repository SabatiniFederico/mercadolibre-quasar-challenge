package service

import "github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"

var sattellites = []model.Satellite{
	{
		Name: "Kenobi",
		Pos:  model.Point{X: -500, Y: -200},
	},
	{
		Name: "Skywalker",
		Pos:  model.Point{X: 100, Y: -100},
	},
	{
		Name: "Sato",
		Pos:  model.Point{X: 500, Y: 100},
	},
}

func GetLocation(distances ...float32) (x, y float32) {
	//var coordinateX float32
	//var coordinateY float32

	return 20, 20
}

func GetMessage(messages ...[]string) (msg string) {
	return "some message"
}
