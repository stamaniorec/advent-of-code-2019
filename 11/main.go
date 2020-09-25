package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"./painting_robot"
	p "./panel"
	v "./vector"
)

func readRobotProgram() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func main() {
	program, err := readRobotProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	robotInput := make(chan p.PanelColor, 1)
	robotColorOutput := make(chan p.PanelColor, 1)
	robotDirectionOutput := make(chan v.Vector, 1)

	go painting_robot.Run(program, robotInput, robotColorOutput, robotDirectionOutput)

	// Part 1
	// paintedPanels := p.GetPaintedPanels(p.Black, robotInput, robotColorOutput, robotDirectionOutput)
	// fmt.Println(len(paintedPanels))

	// Part 2
	paintedPanels := p.GetPaintedPanels(p.White, robotInput, robotColorOutput, robotDirectionOutput)
	grid := p.ConstructGrid(paintedPanels)

	p.PrintRegistrationIdentifier(grid)
}
