package model

type Position struct {
	X float32
	Y float32
}

type Satellite struct {
	Name string
	Pos  Position
}

var Sattellites = []Satellite{
	Satellite{
		Name: "Kenobi",
		Pos:  Position{-500, -200},
	},
	Satellite{
		Name: "Skywalker",
		Pos:  Position{100, -100},
	},
	Satellite{
		Name: "Sato",
		Pos:  Position{500, 100},
	},
}
