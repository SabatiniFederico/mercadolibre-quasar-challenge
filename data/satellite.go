package data

type OnlineSatellite struct {
	Name        string
	Xcoordinate float64
	Ycoordinate float64
}

var onlineSatellites = []OnlineSatellite{
	{
		Name:        "kenobi",
		Xcoordinate: -500.0,
		Ycoordinate: -200.0,
	},
	{
		Name:        "skywalker",
		Xcoordinate: 100.0,
		Ycoordinate: -100.0,
	},
	{
		Name:        "sato",
		Xcoordinate: 500.0,
		Ycoordinate: 100.0,
	},
}

func GetOnlineSatellites() []OnlineSatellite {
	return onlineSatellites
}
