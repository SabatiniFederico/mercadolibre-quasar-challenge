package triangulation

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
	return 20.0, 20.0
}

func GetMessage(messages ...[]string) (msg string) {
	return "some message"
}
