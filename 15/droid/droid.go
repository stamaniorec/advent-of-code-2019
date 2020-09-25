package droid

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

func Run(program string, input <-chan int, output chan<- int) error {
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

	for {
		direction, ok := <-input
		if !ok {
			break
		}
		io.WriteString(stdin, fmt.Sprintf("%d\n", direction))

		ok = scanner.Scan()
		if !ok {
			break
		}

		outputLine := scanner.Text()
		outputAsNumber, _ := strconv.Atoi(strings.TrimSpace(outputLine))

		output <- outputAsNumber
	}

	close(output)

	return cmd.Wait()
}
