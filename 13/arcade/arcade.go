package arcade

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// This executable contains the intcode computer from Day 9
// const IntcodeComputerExecutable = "../09/main"

// This executable contains the free playing mode for part 2
const IntcodeComputerExecutable = "./13_part2"

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
		// Detect if no output is being produced
		// i.e. if the intcode computer process is blocked waiting for user input
		outputTimeout := time.AfterFunc(time.Millisecond*15, func() {
			joystickInput, ok := <-input
			if !ok {
				return
			}

			// Provide the joystick input
			// This will unblock the scanner.Scan() call
			io.WriteString(stdin, fmt.Sprintf("%d\n", joystickInput))
		})

		// This is a blocking operation
		ok := scanner.Scan()
		if !ok {
			break
		}

		// If the scanner has read output, cancel the timeout.
		outputTimeout.Stop()

		outputLine := scanner.Text()
		outputAsNumber, _ := strconv.Atoi(strings.TrimSpace(outputLine))

		output <- outputAsNumber
	}

	close(output)

	return cmd.Wait()
}
