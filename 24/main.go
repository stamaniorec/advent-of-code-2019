package main

import (
	"bufio"
	"fmt"
	"hash/fnv"

	"os"
)

const GridSize = 5

func readGrid() ([][]rune, error) {
	grid := make([][]rune, GridSize)
	for row := range grid {
		grid[row] = make([]rune, GridSize)
	}

	f, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	var row int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		for i, c := range line {
			grid[row][i] = c
		}
		row++
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading input file: %w", err)
	}

	return grid, nil
}

func getGridHash(grid [][]rune) int64 {
	h := fnv.New32a()
	for row := range grid {
		for col := range grid[row] {
			h.Write([]byte{byte(grid[row][col])})
		}
	}

	return int64(h.Sum32())
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func part1(grid [][]rune) {
	fmt.Println("Initial state:")
	printGrid(grid)
	fmt.Println()

	stateMap := make(map[int64]bool)
	stateMap[getGridHash(grid)] = true

	for min := 1; min <= 10000; min++ {
		newGrid := make([][]rune, GridSize)
		for row := range newGrid {
			newGrid[row] = make([]rune, GridSize)
		}

		// fmt.Printf("After %d minute:\n", min)

		for i := 0; i < GridSize; i++ {
			for j := 0; j < GridSize; j++ {
				hasBugAbove := i > 0 && grid[i-1][j] == '#'
				hasBugBelow := i < GridSize-1 && grid[i+1][j] == '#'
				hasBugToTheLeft := j > 0 && grid[i][j-1] == '#'
				hasBugToTheRight := j < GridSize-1 && grid[i][j+1] == '#'
				bugs := []bool{hasBugAbove, hasBugBelow, hasBugToTheLeft, hasBugToTheRight}
				var cntBugs int
				for _, hasBug := range bugs {
					if hasBug {
						cntBugs++
					}
				}

				if grid[i][j] == '.' && (cntBugs == 1 || cntBugs == 2) {
					newGrid[i][j] = '#'
				} else if grid[i][j] == '#' && cntBugs != 1 {
					newGrid[i][j] = '.'
				} else {
					newGrid[i][j] = grid[i][j]
				}
			}
		}

		hash := getGridHash(newGrid)
		if _, found := stateMap[hash]; found {
			var biodiversityRating int64
			var powOf2 int64 = 1
			for i := 0; i < GridSize; i++ {
				for j := 0; j < GridSize; j++ {
					if newGrid[i][j] == '#' {
						biodiversityRating += powOf2
					}
					powOf2 *= 2
				}
			}
			fmt.Println(biodiversityRating)
			return
		}

		stateMap[hash] = true

		// printGrid(newGrid)
		// fmt.Println()

		grid = newGrid
	}
}

type grid [][]rune

func makeGrid() grid {
	g := make([][]rune, GridSize)
	for i := range g {
		g[i] = make([]rune, GridSize)
		for j := range g[i] {
			g[i][j] = '.'
		}
	}
	return grid(g)
}

func part2(initialGrid [][]rune) {
	var gridsDeque []grid
	gridsDeque = append(gridsDeque, initialGrid)

	for min := 1; min <= 200; min++ {
		gridsDeque = append([]grid{makeGrid()}, gridsDeque...)
		gridsDeque = append(gridsDeque, makeGrid())

		var newGridsDeque []grid
		for indexInDeque, g := range gridsDeque {
			newGrid := makeGrid()

			for i := 0; i < GridSize; i++ {
				for j := 0; j < GridSize; j++ {
					above := i > 0 && g[i-1][j] == '#'
					below := i < GridSize-1 && g[i+1][j] == '#'
					left := j > 0 && g[i][j-1] == '#'
					right := j < GridSize-1 && g[i][j+1] == '#'
					adj := []bool{above, below, left, right}

					// look above into next (+1) depth
					if i == 3 && j == 2 && indexInDeque < len(gridsDeque)-1 {
						for _, el := range gridsDeque[indexInDeque+1][GridSize-1] {
							adj = append(adj, el == '#')
						}
					}

					// look below into next (+1) depth
					if i == 1 && j == 2 && indexInDeque < len(gridsDeque)-1 {
						for _, el := range gridsDeque[indexInDeque+1][0] {
							adj = append(adj, el == '#')
						}
					}

					// look to the left into next (+1) depth
					if i == 2 && j == 3 && indexInDeque < len(gridsDeque)-1 {
						for _, row := range gridsDeque[indexInDeque+1] {
							for colIndex, el := range row {
								if colIndex == GridSize-1 {
									adj = append(adj, el == '#')
								}
							}
						}
					}

					// look to the right into next (+1) depth
					if i == 2 && j == 1 && indexInDeque < len(gridsDeque)-1 {
						for _, row := range gridsDeque[indexInDeque+1] {
							for colIndex, el := range row {
								if colIndex == 0 {
									adj = append(adj, el == '#')
								}
							}
						}
					}

					// look above into prev (-1) depth
					if i == 0 && indexInDeque > 0 {
						adj = append(adj, gridsDeque[indexInDeque-1][1][2] == '#')
					}

					// look below into prev (-1) depth
					if i == GridSize-1 && indexInDeque > 0 {
						adj = append(adj, gridsDeque[indexInDeque-1][3][2] == '#')
					}

					// look to the left into prev (-1) depth
					if j == 0 && indexInDeque > 0 {
						adj = append(adj, gridsDeque[indexInDeque-1][2][1] == '#')
					}

					// look to the right into prev (-1) depth
					if j == GridSize-1 && indexInDeque > 0 {
						adj = append(adj, gridsDeque[indexInDeque-1][2][3] == '#')
					}

					var cntBugs int
					for _, hasBug := range adj {
						if hasBug {
							cntBugs++
						}
					}

					if i == 2 && j == 2 && indexInDeque < len(gridsDeque)-1 {
						newGrid[i][j] = '.'
					} else if g[i][j] == '.' && (cntBugs == 1 || cntBugs == 2) {
						newGrid[i][j] = '#'
					} else if g[i][j] == '#' && cntBugs != 1 {
						newGrid[i][j] = '.'
					} else {
						newGrid[i][j] = g[i][j]
					}
				}
			}

			newGridsDeque = append(newGridsDeque, newGrid)
		}

		gridsDeque = newGridsDeque
	}

	var bugsCnt int
	for _, g := range gridsDeque {
		for _, row := range g {
			for _, el := range row {
				if el == '#' {
					bugsCnt++
				}
			}
		}
	}
	fmt.Println(bugsCnt)
}

func main() {
	grid, err := readGrid()
	if err != nil {
		fmt.Println(err)
		return
	}

	// part1(grid)
	part2(grid)
}
