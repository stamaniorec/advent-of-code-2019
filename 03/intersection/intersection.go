package intersection

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	c "../coordinate"
)

type IntersectionResult struct {
	Intersection                c.Coordinate
	FewestStepsFromCentralPoint int
}

func FindIntersections(wire1Path string, wire2Path string) ([]IntersectionResult, error) {
	visitedCoordinatesWithStepsCount := make(map[c.Coordinate]int)
	var intersections []IntersectionResult

	err := WalkWirePath(wire1Path, func(c c.Coordinate, step int) {
		visitedCoordinatesWithStepsCount[c] = step
	})
	if err != nil {
		return make([]IntersectionResult, 0), err
	}

	err = WalkWirePath(wire2Path, func(c c.Coordinate, step int) {
		if otherSteps, isVisited := visitedCoordinatesWithStepsCount[c]; isVisited {
			intersections = append(intersections, IntersectionResult{
				Intersection:                c,
				FewestStepsFromCentralPoint: step + otherSteps,
			})
		}
	})
	if err != nil {
		return make([]IntersectionResult, 0), err
	}

	return intersections, nil
}

func MinManhattanDistanceToIntersection(intersections []IntersectionResult) int {
	centralPort := c.GetCentralPort()

	minDistance := math.MaxInt32
	for _, intersectionResult := range intersections {
		d := c.CalculateManhattanDistance(centralPort, intersectionResult.Intersection)
		if d < minDistance {
			minDistance = d
		}
	}

	return minDistance
}

func FewestStepsFromCentralPointIntersection(intersections []IntersectionResult) int {
	fewestSteps := math.MaxInt32
	for _, intersectionResult := range intersections {
		if intersectionResult.FewestStepsFromCentralPoint < fewestSteps {
			fewestSteps = intersectionResult.FewestStepsFromCentralPoint
		}
	}

	return fewestSteps
}

type onStep func(c c.Coordinate, step int)

func WalkWirePath(wirePath string, onStep onStep) error {
	centralPort := c.GetCentralPort()

	currentPosition := centralPort
	var step int

	instructions := strings.Split(wirePath, ",")
	for _, instruction := range instructions {
		direction := string(instruction[0])
		count, _ := strconv.Atoi(instruction[1:])

		for i := 0; i < count; i++ {
			switch direction {
			case "L":
				currentPosition = c.GoLeft(currentPosition)
			case "R":
				currentPosition = c.GoRight(currentPosition)
			case "U":
				currentPosition = c.GoUp(currentPosition)
			case "D":
				currentPosition = c.GoDown(currentPosition)
			default:
				return fmt.Errorf("Unknown direction %s", direction)
			}

			step++
			onStep(currentPosition, step)
		}
	}

	return nil
}
