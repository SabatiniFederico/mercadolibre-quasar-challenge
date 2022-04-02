package trilateration

import (
	"math"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

func translatePoints(translation model.Point, points ...model.Point) []model.Point {

	for i, point := range points {
		point.X = point.X + translation.X
		point.Y = point.Y + translation.Y
		points[i] = point
	}
	return points
}

func rotatePoints(rotation float64, points ...model.Point) []model.Point {

	for i, point := range points {
		x := (point.X*math.Cos(rotation) + point.Y*-math.Sin(rotation))
		y := (point.X*(math.Sin(rotation)) + point.Y*math.Cos(rotation))
		points[i] = model.Point{X: x, Y: y}
	}

	return points
}
