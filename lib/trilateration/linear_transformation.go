package trilateration

import (
	"math"
)

func linearPointsTranslation(translation Point, points ...Point) []Point {
	var translatedPoints []Point
	for _, point := range points {
		translatedPoints = append(translatedPoints, translatePoint(translation, point))
	}
	return translatedPoints
}

func translatePoint(translation Point, point Point) Point {
	x := point.X + translation.X
	y := point.Y + translation.Y
	return Point{X: x, Y: y}
}

func linearPointsRotation(rotation float64, points ...Point) []Point {
	var rotatedPoints []Point
	for _, point := range points {
		rotatedPoints = append(rotatedPoints, rotatePoint(rotation, point))
	}
	return rotatedPoints
}

func rotatePoint(rotation float64, point Point) Point {
	x := (point.X*math.Cos(rotation) + point.Y*-math.Sin(rotation))
	y := (point.X*(math.Sin(rotation)) + point.Y*math.Cos(rotation))
	return Point{X: x, Y: y}
}
