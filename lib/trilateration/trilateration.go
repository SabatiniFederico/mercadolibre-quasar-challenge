package trilateration

import (
	"errors"
	"math"
)

var ErrNoSolution = errors.New("there is no posible solution")

const marginOfDistanceError = 1

type Point struct {
	X float64
	Y float64
}

func Solve2DTrilateration(points []Point, distances []float64) (Point, error) {

	r1 := distances[0]
	r2 := distances[1]
	r3 := distances[2]

	//translating point1 to origin
	translation := Point{X: -points[0].X, Y: -points[0].Y}
	tPoints := linearPointsTranslation(translation, points...)

	//rotating point2 to X axis.
	rotation := math.Acos(tPoints[1].X / math.Sqrt(math.Pow(tPoints[1].X, 2)+math.Pow(tPoints[1].Y, 2)))
	rotatedPoints := linearPointsRotation(-rotation, tPoints...)

	coordinateX := (math.Pow(r1, 2) - math.Pow(r2, 2) + math.Pow(rotatedPoints[1].X, 2)) / (2 * rotatedPoints[1].X)
	coordinateY := math.Sqrt(math.Pow(r1, 2) - math.Pow(coordinateX, 2))

	//i have to check two possible solutions for Y
	if isAccurateDistance(r3, normalize(rotatedPoints[2].X-coordinateX, rotatedPoints[2].Y+coordinateY)) {
		coordinateY = -coordinateY
	}

	if isAccurateDistance(r3, normalize(rotatedPoints[2].X-coordinateX, rotatedPoints[2].Y-coordinateY)) {
		return translatePoint(points[0], rotatePoint(rotation, Point{X: coordinateX, Y: coordinateY})), nil
	}

	return Point{}, ErrNoSolution
}

func normalize(x float64, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func isAccurateDistance(expectedDistance float64, actualDistance float64) bool {
	return math.Abs(expectedDistance-actualDistance) <= marginOfDistanceError
}
