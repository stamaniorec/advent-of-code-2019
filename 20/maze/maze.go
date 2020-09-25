package maze

import (
	p "./position"
)

type Maze struct {
	Grid          [][]rune
	Width, Height int
	InnerHeight, InnerWidth int
	TopLeft, TopRight, BottomLeft p.Position // todo: leave TopLeft only
	InnerTopLeft p.Position
}

func BuildMaze(grid [][]rune) Maze {
	maze := Maze{
		Grid:   grid,
		Width:  findMazeWidth(grid),
		Height: findMazeHeight(grid),
		TopLeft: findTopLeft(grid),
	}

	maze.TopRight = p.Position{
		Row: maze.TopLeft.Row,
		Col: maze.TopLeft.Col + maze.Width - 1,
	}
	maze.BottomLeft = p.Position{
		Row: maze.TopLeft.Row + maze.Height - 1,
		Col: maze.TopLeft.Col,
	}

	maze.InnerTopLeft = findInnerTopLeft(grid)
	maze.InnerHeight = findMazeInnerHeight(grid)
	maze.InnerWidth = findMazeInnerWidth(grid)

	return maze
}



func findTopMazeRowIndex(grid [][]rune) int {
	midCol := len(grid[0]) / 2

	for i, row := range grid {
		if row[midCol] == '.' || row[midCol] == '#' {
			return i
		}
	}

	return -1
}

func findTopLeft(grid [][]rune) p.Position {
	row := findTopMazeRowIndex(grid)
  for i, c := range grid[row] {
		if c == '.' || c == '#' {
			return p.Position{
				Row: row,
				Col: i,
			}
		}
	}

	return p.Position{}
}

func findInnerTopLeft(grid [][]rune) p.Position {
	midRow := len(grid) / 2
	midCol := len(grid[0]) / 2

	var rowIndex int
	for rowIndex == 0 {
		if grid[midRow][midCol] == '.' || grid[midRow][midCol] == '#' {
			rowIndex = midRow
		}
		midRow--
	}

	var colIndex int
	for colIndex == 0 {
		if grid[rowIndex+1][midCol] == '.' || grid[rowIndex+1][midCol] == '#' {
			colIndex = midCol
		}
		midCol--
	}

	return p.Position{
		Row: rowIndex,
		Col: colIndex,
	}
}

func findMazeWidth(grid [][]rune) int {
	var width int
	topRow := findTopMazeRowIndex(grid)

	for _, col := range grid[topRow] {
		if col == '.' || col == '#' {
			width++
		}
	}
	return width
}

func findMazeHeight(grid [][]rune) int {
	var height int
	TopLeft := findTopLeft(grid)

	for _, row := range grid {
		if len(row) < TopLeft.Col {
			continue
		}

		if row[TopLeft.Col] == '.' || row[TopLeft.Col] == '#' {
			height++
		}
	}
	return height
}

func findMazeInnerHeight(grid [][]rune) int {
	var height int
	TopLeft := findInnerTopLeft(grid)

	TopLeft.Row++
	for grid[TopLeft.Row][TopLeft.Col+1] != '#' && grid[TopLeft.Row][TopLeft.Col+1] != '.'{
		height++
		TopLeft.Row++
	}
	return height
}

func findMazeInnerWidth(grid [][]rune) int {
	var width int
	TopLeft := findInnerTopLeft(grid)

	TopLeft.Col++
	for grid[TopLeft.Row+1][TopLeft.Col] != '#' && grid[TopLeft.Row+1][TopLeft.Col] != '.'{
		width++
		TopLeft.Col++
	}
	return width
}
