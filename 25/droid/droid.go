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

func Run(program string, input <-chan string) (error) {
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
	defer cmd.Wait()

	var outputBuffer []rune

	for scanner.Scan() {
		outputLine := scanner.Text()
		outputLineAsNumber, _ := strconv.Atoi(outputLine)
		curChar := rune(outputLineAsNumber)
		fmt.Print(string(curChar))

		outputBuffer = append(outputBuffer, curChar)
		if len(outputBuffer) > len("Command?") {
			outputBuffer = outputBuffer[1:len(outputBuffer)]
		}

		if string(outputBuffer) == "Command?" {
			outputBuffer = make([]rune, 0)
			fmt.Println()

			command := <-input
			for _, c := range command {
				io.WriteString(stdin, fmt.Sprintf("%d\n", int(c)))
			}
			io.WriteString(stdin, "10\n") // newline
		}
	}

	return nil
}
