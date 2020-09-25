package direction

import (
	p "../maze/position"
)

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

var DirectionVectors = map[Direction]p.Position{
	Right: p.Position{0, 1},
	Down:  p.Position{1, 0},
	Left:  p.Position{0, -1},
	Up:    p.Position{-1, 0},
}
