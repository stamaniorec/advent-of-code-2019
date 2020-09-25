package grid

import (
	"fmt"
	"hash/fnv"
	"strings"
)

type Position struct {
	Row, Col int
}

func BuildGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]rune, rows)
	for row := range grid {
		grid[row] = make([]rune, cols)
	}

	var col, row int
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			row++
			col = 0
			continue
		}

		grid[row][col] = rune(input[i])
		col++
	}

	return grid
}

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func GetGridHash(grid [][]rune) int64 {
	h := fnv.New32a()
	for row := range grid {
		for col := range grid[row] {
			h.Write([]byte{byte(grid[row][col])})
		}
	}

	return int64(h.Sum32())
}

func FindDoors(grid [][]rune) map[rune]Position {
	doors := make(map[rune]Position)
	for i, row := range grid {
		for j, cell := range row {
			if isDoor := cell >= 'A' && cell <= 'Z'; isDoor {
				doors[cell] = Position{Row: i, Col: j}
			}
		}
	}
	return doors
}

func CountKeys(grid [][]rune) int {
	var count int
	for _, row := range grid {
		for _, cell := range row {
			if isKey := cell >= 'a' && cell <= 'z'; isKey {
				count++
			}
		}
	}
	return count
}

func FindRobotPositions(grid [][]rune) []Position {
	var robotPositions []Position
	for i, row := range grid {
		for j, cell := range row {
			if isRobotPosition := cell == '@'; isRobotPosition {
				robotPositions = append(robotPositions, Position{
					Row: i,
					Col: j,
				})
			}
		}
	}
	return robotPositions
}

func ObtainKey(grid [][]rune, keyPos, robotPos Position, doors map[rune]Position) {
	keyLabel := grid[keyPos.Row][keyPos.Col]

	// Move robot from robotPos to keyPos
	grid[keyPos.Row][keyPos.Col] = '@'
	grid[robotPos.Row][robotPos.Col] = '.'

	// Unlock the door
	doorLabel := keyLabel - 'a' + 'A'
	if door, found := doors[doorLabel]; found {
		grid[door.Row][door.Col] = '.'
	}
}

func PutBackKey(grid [][]rune, keyPos, origPos Position, label rune, doors map[rune]Position) {
	// Move robot to origPos back from keyPos
	grid[keyPos.Row][keyPos.Col] = label
	grid[origPos.Row][origPos.Col] = '@'

	// Put back the door
	doorLabel := label - 'a' + 'A'
	if door, found := doors[doorLabel]; found {
		grid[door.Row][door.Col] = doorLabel
	}
}
