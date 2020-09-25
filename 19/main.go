package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"./drone"
)

func readProgram() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func countPoundsInRow(program string, row, colsCount int) int {
	var cnt int
	var foundPound bool
	for col := 0; col < colsCount; col++ {
		res, err := drone.Run(program, row, col)
		if err != nil {
			panic("Running droid program failed")
		}

		isPound := res == 1
		if isPound {
			cnt++
			foundPound = true
		}

		if !isPound && foundPound {
			// The pattern is [(dots), (pounds), (dots)]
			// so if you find a dot and you've already seen a pound
			// then you're in the second (dots), so no need to continue
			// searching for pounds
			break
		}
	}
	return cnt
}

// findFirstRowWithNPounds does binary search to find which row to start from
func findFirstRowWithNPounds(program string, n, colsCount int) int {
	var i int
	for step := colsCount + 1; step >= 1; step /= 2 {
		for {
			makeStepIndex := i + step
			if countPoundsInRow(program, makeStepIndex, colsCount) >= n {
				break
			}

			i = makeStepIndex
		}
	}
	return i + 1
}

func main() {
	program, err := readProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Part 1
	// colsCount := 50

	// var cnt int
	// for i := 0; i < colsCount; i++ {
	// 	fmt.Printf("at row %d\n", i)
	// 	cnt += countPoundsInRow(program, i, colsCount)
	// }
	// fmt.Println(cnt)

	// Part 2
	maxCols := 800 // choose some "large enough" number
	squareSize := 100

	firstRow := findFirstRowWithNPounds(program, squareSize, maxCols)
	fmt.Println(firstRow)

	row := firstRow
	var found bool
	for !found {
		for col := 0; col < maxCols; col++ {
			topLeft, err := drone.Run(program, row, col)
			if err != nil {
				fmt.Println(err)
				return
			}

			if isPound := topLeft == 1; !isPound {
				continue
			}

			topRight, err := drone.Run(program, row, col+squareSize-1)
			if err != nil {
				fmt.Println(err)
				return
			}

			if isPound := topRight == 1; !isPound {
				break
			}

			bottomLeft, err := drone.Run(program, row+squareSize-1, col)
			if err != nil {
				fmt.Println(err)
				return
			}

			if isPound := bottomLeft == 1; !isPound {
				continue
			}

			bottomRight, err := drone.Run(program, row+squareSize-1, col+squareSize-1)
			if err != nil {
				fmt.Println(err)
				return
			}

			if isPound := bottomRight == 1; !isPound {
				continue
			}

			found = true
			fmt.Printf("%d,%d\n", row, col)
		}
		row++
	}

	// 1011,555
}
