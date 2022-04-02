package trilateration

import (
	"math"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/model"
)

func translatePoints(translation model.Point, points ...model.Point) []model.Point {

	for i, point := range points {
		points[i] = translatePoint(translation, point)
	}
	return points
}

func translatePoint(translation model.Point, point model.Point) model.Point {
	x := point.X + translation.X
	y := point.Y + translation.Y
	return model.Point{X: x, Y: y}
}

func rotatePoints(rotation float64, points ...model.Point) []model.Point {

	for i, point := range points {
		points[i] = rotatePoint(rotation, point)
	}
	return points
}

func rotatePoint(rotation float64, point model.Point) model.Point {
	x := (point.X*math.Cos(rotation) + point.Y*-math.Sin(rotation))
	y := (point.X*(math.Sin(rotation)) + point.Y*math.Cos(rotation))
	return model.Point{X: x, Y: y}
}
