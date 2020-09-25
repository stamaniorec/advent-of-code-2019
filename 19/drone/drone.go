package drone

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// This executable contains the intcode computer from Day 9
const IntcodeComputerExecutable = "../09/main"

func Run(program string, row, col int) (int, error) {
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

	io.WriteString(stdin, fmt.Sprintf("%d\n", row))
	io.WriteString(stdin, fmt.Sprintf("%d\n", col))

	ok := scanner.Scan()
	if !ok {
		return 0, nil
	}

	outputLine := scanner.Text()
	outputAsNumber, _ := strconv.Atoi(strings.TrimSpace(outputLine))

	return outputAsNumber, nil
}
