package entity

type ClassifiedMessage struct {
	Name     string   `json:"name" validate:"required,eq=kenobi|eq=skywalker|eq=sato"`
	Distance float64  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required"`
}

type SatellitesRequest struct {
	Satellites []ClassifiedMessage `json:"satellites"`
}
