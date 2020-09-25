package droid

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"os/exec"
)

// This executable contains the intcode computer from Day 9
const IntcodeComputerExecutable = "../09/main"

func Run(program string, instructions []string) (int, error) {
	cmd := exec.Command(IntcodeComputerExecutable, program)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	if err != nil {
		return 0, err
	}
	defer cmd.Wait()

	for _, instr := range instructions {
		for _, c := range instr {
			io.WriteString(stdin, fmt.Sprintf("%d\n", int(c)))
		}
		io.WriteString(stdin, "10\n") // newline
	}

	var res int
	for scanner.Scan() {
		outputLine := scanner.Text()
		outputLineAsNumber, _ := strconv.Atoi(outputLine)
		res = outputLineAsNumber
		fmt.Print(string(rune(outputLineAsNumber)))
	}
	fmt.Println()

	return res, nil
}
