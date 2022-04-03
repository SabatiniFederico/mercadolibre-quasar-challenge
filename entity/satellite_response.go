package entity

type SolutionResponse struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}
