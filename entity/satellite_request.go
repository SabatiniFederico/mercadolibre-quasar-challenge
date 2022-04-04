package entity

type Satellite struct {
	Name     string   `json:"name" validate:"required,eq=kenobi|eq=skywalker|eq=sato"`
	Distance float64  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required"`
}

type SatellitesRequest struct {
	Satellites []Satellite `json:"satellites" validate:"dive,required"`
}
