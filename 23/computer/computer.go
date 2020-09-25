package computer

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	// "sync"

	"os/exec"
)

// This executable contains the intcode computer from Day 9
const IntcodeComputerExecutable = "../09/main"

func Run(program string, address int, input <-chan int, output chan<- int) (error) {
	cmd := exec.Command(IntcodeComputerExecutable, program)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewReader(stdout)

	err = cmd.Start()
	if err != nil {
		return err
	}
	defer cmd.Wait()

	io.WriteString(stdin, fmt.Sprintf("%d\n", address))

	go func(scanner *bufio.Reader, address int) {
		for {
			outputLine, err := scanner.ReadString('\n')

			if err == nil {
				outputLineAsNumber, err := strconv.Atoi(strings.TrimSpace(outputLine))
				if err != nil {
					panic(err)
				}
				output <- outputLineAsNumber
				// fmt.Printf("%d wrote  %d\n", address, outputLineAsNumber)
			}
		}
	}(scanner, address)

	for {
		select {
			case x := <-input:
				io.WriteString(stdin, fmt.Sprintf("%d\n", x))

				if x != -1 {
					select {
					case y := <-input:
						io.WriteString(stdin, fmt.Sprintf("%d\n", y))
					default:
						panic("Not good. Run again.")
					}
				}
			default:
				io.WriteString(stdin, fmt.Sprintf("%d\n", -1))
		}
	}

	return nil
}
