package painting_robot

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"

	p "../panel"
	v "../vector"
)

const IntcodeComputerExecutable = "../09/main"

var directionVectors = []v.Vector{
	v.Vector{X: 0, Y: -1},
	v.Vector{X: 1, Y: 0},
	v.Vector{X: 0, Y: 1},
	v.Vector{X: -1, Y: 0},
}

func Run(program string,
	colorInput <-chan p.PanelColor,
	colorOutput chan<- p.PanelColor,
	directionOutput chan<- v.Vector) error {

	cmd := exec.Command(IntcodeComputerExecutable, program)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	if err != nil {
		return err
	}

	var directionVectorIndex int

	for {
		currentColor := <-colorInput
		io.WriteString(stdin, fmt.Sprintf("%d\n", currentColor))

		if ok := scanner.Scan(); !ok {
			break
		}
		newColor := getNewColor(scanner.Text())
		colorOutput <- newColor

		if ok := scanner.Scan(); !ok {
			break
		}
		directionCode := getDirectionCode(scanner.Text())
		directionVectorIndex = getDirectionVectorIndex(directionVectorIndex, directionCode)
		directionOutput <- directionVectors[directionVectorIndex]
	}

	close(colorOutput)
	close(directionOutput)

	return cmd.Wait()
}

func getNewColor(outputLine string) p.PanelColor {
	newColor, _ := strconv.Atoi(strings.TrimSpace(outputLine))
	return p.PanelColor(newColor)
}

func getDirectionCode(outputLine string) int {
	directionCode, _ := strconv.Atoi(strings.TrimSpace(outputLine))
	return directionCode
}

func getDirectionVectorIndex(currentIndex, directionCode int) int {
	if directionCode == 0 {
		return rotate90DegreesLeft(currentIndex)
	}

	return rotate90DegreesRight(currentIndex)
}

func rotate90DegreesLeft(index int) int {
	index--
	if index < 0 {
		index = len(directionVectors) - 1
	}

	return index
}

func rotate90DegreesRight(index int) int {
	index++
	if index >= len(directionVectors) {
		index = 0
	}

	return index
}
