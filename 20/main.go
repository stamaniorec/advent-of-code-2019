package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	d "./direction"
	m "./maze"
	p "./maze/position"
	prt "./portal"
)

func readInput() (string, error) {
	dat, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return string(dat), nil
}

// BuildGrid uses the content of the input file
// and transforms it into a rune matrix.
// It pads the matrix with extra rows and columns
// so that later on out of bounds checks can be skipped
func BuildGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[5]) + 2

	grid := make([][]rune, rows+2)

	// Pad a row on the top
	grid[0] = make([]rune, cols)

	for i, line := range lines {
		grid[i+1] = make([]rune, cols+2)

		// Pad a column on the left
		grid[i+1][0] = ' '

		copy(grid[i+1][1:cols+1], []rune(line))

		// Pad a column on the right
		grid[i+1][cols+1] = ' '
	}

	// Pad a row on the bottom
	grid[rows] = make([]rune, cols)

	return grid
}

func bfs(maze m.Maze, startingPos, endingPos p.Position, portals []prt.Portal) int {
	type queueItem struct {
		pos   p.Position
		steps int
	}

	var q []queueItem
	q = append(q, queueItem{
		pos:   startingPos,
		steps: 0,
	})

	visited := make(map[p.Position]bool)
	visited[startingPos] = true

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		visited[item.pos] = true

		var next []p.Position
		for _, dv := range d.DirectionVectors {
			next = append(next, p.Position{
				Row: item.pos.Row + dv.Row,
				Col: item.pos.Col + dv.Col,
			})
		}

		for _, n := range next {
			if visited[n] {
				continue
			}

			if n == endingPos {
				return item.steps + 1
			}

			if ch := maze.Grid[n.Row][n.Col]; ch == '.' {
				portal := prt.GetPortalAt(portals, n)
				defaultPortal := prt.Portal{}
				if portal != defaultPortal {
					opposite := prt.FindOppositePortal(portals, portal)
					fmt.Println(opposite.Label)
					fmt.Println(n)
					fmt.Println(opposite.Position)
					n = opposite.Position
					item.steps++
				}

				q = append(q, queueItem{
					pos:   n,
					steps: item.steps + 1,
				})
			}
		}
	}

	return 0
}

func bfsPart2(maze m.Maze, startingPos, endingPos p.Position, innerPortals, outerPortals []prt.Portal) int {
	var portals []prt.Portal
	portals = append(portals, outerPortals...)
	portals = append(portals, innerPortals...)

	type queueItem struct {
		pos   p.Position
		steps int
		level int
	}
	type visitedItem struct {
		pos   p.Position
		level int
	}

	var q []queueItem
	q = append(q, queueItem{
		pos:   startingPos,
		steps: 0,
		level: 0,
	})

	visited := make(map[visitedItem]bool)
	visited[visitedItem{startingPos, 0}] = true

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		visited[visitedItem{item.pos, item.level}] = true

		var next []p.Position
		for _, dv := range d.DirectionVectors {
			next = append(next, p.Position{
				Row: item.pos.Row + dv.Row,
				Col: item.pos.Col + dv.Col,
			})
		}

		for _, n := range next {
			if visited[visitedItem{n, item.level}] {
				continue
			}

			if n == endingPos && item.level == 0 {
				return item.steps + 1
			}

			isWithinBounds :=
				n.Row >= 0 && n.Col >= 0 &&
					n.Row < len(maze.Grid) && n.Col < len(maze.Grid[10])
			if !isWithinBounds {
				continue
			}

			if ch := maze.Grid[n.Row][n.Col]; ch == '.' {
				portal := prt.GetPortalAt(portals, n)
				defaultPortal := prt.Portal{}

				lvl := item.level
				steps := item.steps

				if portalExistsHere := portal != defaultPortal; portalExistsHere {
					opposite := prt.FindOppositePortal(portals, portal)

					var isInnerPortal bool
					for _, x := range innerPortals {
						if x == opposite {
							isInnerPortal = true
							break
						}
					}

					if isInnerPortal && item.level == 0 {
						continue
					}

					n = opposite.Position
					steps++

					if isInnerPortal {
						lvl++
					} else {
						lvl--
					}
				}

				q = append(q, queueItem{
					pos:   n,
					steps: steps + 1,
					level: lvl,
				})
			}
		}
	}

	return 0
}

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := BuildGrid(input)
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}

	maze := m.BuildMaze(grid)

	outerRingPortals := prt.FindOuterRingPortals(maze)
	innerRingPortals := prt.FindInnerRingPortals(maze)

	aaPortal := prt.GetPortal(outerRingPortals, "AA")
	startingPos := aaPortal.Position

	zzPortal := prt.GetPortal(outerRingPortals, "ZZ")
	endingPos := zzPortal.Position

	var portals []prt.Portal
	portals = append(portals, outerRingPortals...)
	portals = append(portals, innerRingPortals...)

	for _, p := range portals {
		fmt.Println(p.Label)
		fmt.Println(p.Position)
	}

	res := bfsPart2(maze, startingPos, endingPos, innerRingPortals, outerRingPortals)
	fmt.Println(res)
}
