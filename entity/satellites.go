package entity

type Point struct {
	X float64
	Y float64
}

type Satellite struct {
	Name string
	Pos  Point
}

type ClassifiedMessage struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatellitesInfo struct {
	Satellites [3]ClassifiedMessage `json:"satellites"`
}

type Solution struct {
	Position Point
	Message  string
}
