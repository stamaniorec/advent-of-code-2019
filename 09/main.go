package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./add"
	"./adjust_relative_base"
	"./equals"
	"./input"
	"./jump_if_false"
	"./jump_if_true"
	"./less_than"
	"./multiply"
	"./output"
)

const inputFileName string = "input.txt"

func ExecuteIntcodeProgram(memory []int) error {
	if len(memory) == 0 {
		return errors.New("Empty input")
	}

	var relativeBase int

	var ip int
	for ip < len(memory) {
		instruction := memory[ip]
		opcode := instruction % 100

		// fmt.Printf("at index %d with instruction %d and opcode %d\n", ip, instruction, opcode)
		// fmt.Printf("relative base is %d\n", relativeBase)

		switch opcode {
		case 1:
			addInstruction := add.Instruction(instruction)
			addInstruction.Execute(memory, ip, relativeBase)

			ip += addInstruction.GetParametersCount() + 1
		case 2:
			multiplyInstruction := multiply.Instruction(instruction)
			multiplyInstruction.Execute(memory, ip, relativeBase)

			ip += multiplyInstruction.GetParametersCount() + 1
		case 3:
			inputInstruction := input.Instruction(instruction)
			inputInstruction.Execute(memory, ip, relativeBase)

			ip += inputInstruction.GetParametersCount() + 1
		case 4:
			outputInstruction := output.Instruction(instruction)
			outputInstruction.Execute(memory, ip, relativeBase)

			ip += outputInstruction.GetParametersCount() + 1
		case 5:
			jumpIfTrueInstruction := jump_if_true.Instruction(instruction)
			newIp := jumpIfTrueInstruction.Execute(memory, ip, relativeBase)

			ip = newIp
		case 6:
			jumpIfFalseInstruction := jump_if_false.Instruction(instruction)
			newIp := jumpIfFalseInstruction.Execute(memory, ip, relativeBase)

			ip = newIp
		case 7:
			lessThanInstruction := less_than.Instruction(instruction)
			lessThanInstruction.Execute(memory, ip, relativeBase)

			ip += lessThanInstruction.GetParametersCount() + 1
		case 8:
			equalsInstruction := equals.Instruction(instruction)
			equalsInstruction.Execute(memory, ip, relativeBase)

			ip += equalsInstruction.GetParametersCount() + 1
		case 9:
			adjustRelativeBaseInstruction := adjust_relative_base.Instruction(instruction)
			relativeBase = adjustRelativeBaseInstruction.Execute(memory, ip, relativeBase)

			ip += adjustRelativeBaseInstruction.GetParametersCount() + 1
		case 99:
			return nil
		default:
			return errors.New(fmt.Sprintf("Unknown opcode %d", opcode))
		}
	}

	return nil
}

func readProgramFromCommandLineArgs() ([]int, error) {
	var program []int
	if len(os.Args) <= 1 {
		return program, errors.New("No program given at standard input")
	}

	splitByComma := strings.Split(os.Args[1], ",")
	for _, el := range splitByComma {
		convEl, _ := strconv.Atoi(el)
		program = append(program, convEl)
	}

	return program, nil
}

func readProgramFromFile() ([]int, error) {
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

const MemorySize = 1000000

func main() {
	var (
		program []int
		err     error
	)

	if len(os.Args) > 1 {
		program, err = readProgramFromCommandLineArgs()
	} else {
		program, err = readProgramFromFile()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	memory := make([]int, MemorySize)
	copy(memory, program)

	err = ExecuteIntcodeProgram(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
