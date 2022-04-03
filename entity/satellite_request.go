package entity

type ClassifiedMessage struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatellitesRequest struct {
	Satellites [3]ClassifiedMessage `json:"satellites"`
}
