package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFileName string = "input.txt"

func ExecuteIntcodeProgram(input []int) error {
	if len(input) == 0 {
		return errors.New("Empty input")
	}

	for i := 0; i < len(input); i += 4 {
		if i+4 >= len(input) {
			return nil
		}

		switch opcode := input[i]; opcode {
		case 1:
			leftOperandIndex := input[i+1]
			rightOperandIndex := input[i+2]
			resultIndex := input[i+3]

			input[resultIndex] = input[leftOperandIndex] + input[rightOperandIndex]
		case 2:
			leftOperandIndex := input[i+1]
			rightOperandIndex := input[i+2]
			resultIndex := input[i+3]

			input[resultIndex] = input[leftOperandIndex] * input[rightOperandIndex]
		case 99:
			break
		default:
			return errors.New(fmt.Sprintf("Unknown opcode %d", opcode))
		}
	}

	return nil
}

func readIntcodeProgramFromInput() ([]int, error) {
	f, err := os.Open(inputFileName)
	if err != nil {
		return nil, fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	var masses []int

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			splitByComma := strings.Split(line, ",")
			var result []int
			for _, el := range splitByComma {
				convEl, _ := strconv.Atoi(el)
				result = append(result, convEl)
			}
			return result, nil
		}
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading input file: %w", err)
	}

	return masses, nil
}

func main() {
	inputFromFile, err := readIntcodeProgramFromInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input := make([]int, len(inputFromFile))
			copy(input, inputFromFile)

			input[1] = i
			input[2] = j

			err = ExecuteIntcodeProgram(input)
			if err != nil {
				fmt.Println(err)
				return
			}

			result := input[0]

			if result == 19690720 {
				fmt.Println(i)
				fmt.Println(j)
			}
		}
	}
}
