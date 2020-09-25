package panel

import (
	"fmt"
	"math"

	v "../vector"
)

type PanelColor int

const (
	Black PanelColor = iota
	White
)

func GetPaintedPanels(startingPanelColor PanelColor,
	robotInput chan<- PanelColor,
	robotColorOutput <-chan PanelColor,
	robotDirectionOutput <-chan v.Vector) map[v.Vector]PanelColor {

	paintedPanels := make(map[v.Vector]PanelColor)

	currentPosition := v.Vector{X: 2, Y: 2}

	paintedPanels[currentPosition] = startingPanelColor

	for {
		currentColor := paintedPanels[currentPosition]
		robotInput <- currentColor

		newColor, ok := <-robotColorOutput
		if !ok {
			break
		}
		paintedPanels[currentPosition] = newColor

		directionVector, ok := <-robotDirectionOutput
		if !ok {
			break
		}

		currentPosition = v.AddVectors(currentPosition, directionVector)
	}

	close(robotInput)

	return paintedPanels
}

func ConstructGrid(paintedPanels map[v.Vector]PanelColor) [][]PanelColor {
	maxX := math.MinInt32
	maxY := math.MinInt32
	for k := range paintedPanels {
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}

	grid := make([][]PanelColor, maxY+1)
	for row := range grid {
		grid[row] = make([]PanelColor, maxX+1)
	}

	for pos, color := range paintedPanels {
		grid[pos.Y][pos.X] = color
	}

	return grid
}

func PrintRegistrationIdentifier(grid [][]PanelColor) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			printPanel(grid[row][col])
		}

		fmt.Println()
	}
}

func printPanel(color PanelColor) {
	if color == Black {
		fmt.Print(".  ")
	} else {
		fmt.Print("#  ")
	}
}
