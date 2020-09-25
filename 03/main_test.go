package main

import (
	"testing"

	"./intersection"
	"github.com/stretchr/testify/assert"
)

func TestMinManhattanDistanceToIntersection(t *testing.T) {
	tests := []struct {
		name      string
		wire1Path string
		wire2Path string
		result    int
	}{
		{
			name:      "t1",
			wire1Path: "R8,U5,L5,D3",
			wire2Path: "U7,R6,D4,L4",
			result:    6,
		},
		{
			name:      "t2",
			wire1Path: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2Path: "U62,R66,U55,R34,D71,R55,D58,R83",
			result:    159,
		},
		{
			name:      "t3",
			wire1Path: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R5",
			wire2Path: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			result:    135,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intersections, err := intersection.FindIntersections(tt.wire1Path, tt.wire2Path)
			assert.Nil(t, err)

			manhattanDistance := intersection.MinManhattanDistanceToIntersection(intersections)
			assert.Equal(t, tt.result, manhattanDistance)
		})
	}
}

func TestFewestStepsFromCentralPointIntersection(t *testing.T) {
	tests := []struct {
		name      string
		wire1Path string
		wire2Path string
		result    int
	}{
		{
			name:      "t1",
			wire1Path: "R8,U5,L5,D3",
			wire2Path: "U7,R6,D4,L4",
			result:    30,
		},
		{
			name:      "t2",
			wire1Path: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2Path: "U62,R66,U55,R34,D71,R55,D58,R83",
			result:    610,
		},
		{
			name:      "t3",
			wire1Path: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R5",
			wire2Path: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			result:    410,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intersections, err := intersection.FindIntersections(tt.wire1Path, tt.wire2Path)
			assert.Nil(t, err)

			fewestSteps := intersection.FewestStepsFromCentralPointIntersection(intersections)
			assert.Equal(t, tt.result, fewestSteps)
		})
	}
}
