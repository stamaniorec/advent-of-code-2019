package camera

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// This executable contains the intcode computer from Day 9
const IntcodeComputerExecutable = "../09/main"

// Viewport represents the output of the Intcode program as a grid
type Viewport struct {
	Grid       [][]int
	Rows, Cols int
}

// ReadViewport runs the Intcode program, which produces camera signals
// and builds a viewport out of them
func ReadViewport(program string) (Viewport, error) {
	signals, err := readCameraSignals(program)
	if err != nil {
		return Viewport{}, err
	}

	return convertCameraSignalsToViewport(signals), nil
}

func PrintViewport(view Viewport) {
	for _, row := range view.Grid {
		for _, cell := range row {
			asString := string(cell)
			fmt.Print(asString)
		}
		fmt.Println()
	}
}

func IsScaffoldIntersection(viewport Viewport, i, j int) bool {
	isScaffold := string(viewport.Grid[i][j]) == "#"
	return isScaffold &&
		i > 0 && j > 0 && i < len(viewport.Grid)-1 && j < len(viewport.Grid[i])-1 &&
		viewport.Grid[i][j] == viewport.Grid[i-1][j] &&
		viewport.Grid[i][j] == viewport.Grid[i][j-1] &&
		viewport.Grid[i][j] == viewport.Grid[i][j+1] &&
		viewport.Grid[i][j] == viewport.Grid[i+1][j]
}

func readCameraSignals(program string) ([]int, error) {
	cmd := exec.Command(IntcodeComputerExecutable, program)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	var signals []int
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}

		outputLine := scanner.Text()
		signal, _ := strconv.Atoi(strings.TrimSpace(outputLine))

		signals = append(signals, signal)
	}

	return signals, cmd.Wait()
}

func convertCameraSignalsToViewport(signals []int) Viewport {
	viewport := Viewport{}
	viewport.Rows = getNumRows(signals)
	viewport.Cols = getNumColumns(signals, viewport.Rows)

	viewport.Grid = make([][]int, viewport.Rows)
	for i := range viewport.Grid {
		viewport.Grid[i] = make([]int, viewport.Cols)
	}

	var row, col int
	for _, x := range signals {
		if isNewline := x == 10; isNewline {
			row++
			col = 0
			continue
		}

		viewport.Grid[row][col] = x
		col++
	}

	return viewport
}

func getNumRows(signals []int) int {
	var rows int
	for _, x := range signals {
		if isNewline := x == 10; isNewline {
			rows++
		}
	}
	return rows - 1
}

func getNumColumns(signals []int, rows int) int {
	return (len(signals) - rows) / rows
}
