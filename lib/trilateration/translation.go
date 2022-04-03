package trilateration

import (
	"math"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

func translatePoints(translation entity.Point, points ...entity.Point) []entity.Point {
	var translatedPoints []entity.Point
	for _, point := range points {
		translatedPoints = append(translatedPoints, translatePoint(translation, point))
	}
	return translatedPoints
}

func translatePoint(translation entity.Point, point entity.Point) entity.Point {
	x := point.X + translation.X
	y := point.Y + translation.Y
	return entity.Point{X: x, Y: y}
}

func rotatePoints(rotation float64, points ...entity.Point) []entity.Point {
	var rotatedPoints []entity.Point
	for _, point := range points {
		rotatedPoints = append(rotatedPoints, rotatePoint(rotation, point))
	}
	return rotatedPoints
}

func rotatePoint(rotation float64, point entity.Point) entity.Point {
	x := (point.X*math.Cos(rotation) + point.Y*-math.Sin(rotation))
	y := (point.X*(math.Sin(rotation)) + point.Y*math.Cos(rotation))
	return entity.Point{X: x, Y: y}
}
