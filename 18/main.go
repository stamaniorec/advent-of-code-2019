package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	g "./grid"
	"./solution"
)

func readInput(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func main() {
	input, err := readInput("input_part2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := g.BuildGrid(input)
	g.PrintGrid(grid)

	robotPositions := g.FindRobotPositions(grid)
	remainingKeys := g.CountKeys(grid)
	doors := g.FindDoors(grid)
	dp := make(map[int64]int)

	minSteps := solution.FindMinSteps(grid, robotPositions, remainingKeys, doors, dp)
	fmt.Println(minSteps)
}
