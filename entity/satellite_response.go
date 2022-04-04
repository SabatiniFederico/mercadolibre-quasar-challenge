package entity

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type StarshipResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
