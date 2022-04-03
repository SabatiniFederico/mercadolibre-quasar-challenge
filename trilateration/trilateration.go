package trilateration

import (
	"errors"
	"math"

	"github.com/SabatiniFederico/mercadolibre-quasar-challenge/entity"
)

var ErrNoSolution = errors.New("there is no posible solution")

const marginOfError = 5

func Solve2DTrilateration(p1 entity.Point, p2 entity.Point, p3 entity.Point, distances []float64) (entity.Point, error) {

	//translating point1 to origin
	points := translatePoints(entity.Point{X: -p1.X, Y: -p1.Y}, p1, p2, p3)

	//rotating point2 to X axis.
	rotation := math.Acos(points[1].X / math.Sqrt(math.Pow(points[1].X, 2)+math.Pow(points[1].Y, 2)))
	rotatedPoints := rotatePoints(-rotation, points...)

	coordinateX := (math.Pow(distances[0], 2) - math.Pow(distances[1], 2) + math.Pow(rotatedPoints[1].X, 2)) / (2 * rotatedPoints[1].X)
	coordinateY := math.Sqrt(math.Pow(distances[0], 2) - math.Pow(coordinateX, 2))

	//i have to check two possible solutions for Y
	if isAccurateDistance(distances[2], normalize(rotatedPoints[2].X-coordinateX, rotatedPoints[2].Y-coordinateY)) {
		return translatePoint(p1, rotatePoint(rotation, entity.Point{X: coordinateX, Y: coordinateY})), nil
	}
	if isAccurateDistance(distances[2], normalize(rotatedPoints[2].X-coordinateX, rotatedPoints[2].Y+coordinateY)) {
		return translatePoint(p1, rotatePoint(rotation, entity.Point{X: coordinateX, Y: coordinateY})), nil
	}

	return entity.Point{}, ErrNoSolution
}

func normalize(x float64, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func isAccurateDistance(expectedDistance float64, actualDistance float64) bool {
	return math.Abs(expectedDistance-actualDistance) <= marginOfError
}
