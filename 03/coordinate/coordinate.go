package coordinate

import (
	"fmt"
	"math"
)

type Coordinate struct {
	Row int
	Col int
}

func GoLeft(c Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row,
		Col: c.Col - 1,
	}
}

func GoRight(c Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row,
		Col: c.Col + 1,
	}
}

func GoUp(c Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row + 1,
		Col: c.Col,
	}
}

func GoDown(c Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row - 1,
		Col: c.Col,
	}
}

func GetCentralPort() Coordinate {
	return Coordinate{
		Row: 1,
		Col: 1,
	}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.Row, c.Col)
}

func CalculateManhattanDistance(c1, c2 Coordinate) int {
	return int(math.Abs(float64(c1.Row-c2.Row)) + math.Abs(float64(c1.Col-c2.Col)))
}
