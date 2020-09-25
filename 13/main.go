package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"./arcade"
)

func readProgram() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func calcJoystickPosition(ball_x, paddle_x int) int {
	if ball_x > paddle_x {
		return 1
	} else if ball_x < paddle_x {
		return -1
	} else {
		return 0
	}
}

func main() {
	program, err := readProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputChannel := make(chan int)
	outputChannel := make(chan int)

	go arcade.Run(program, inputChannel, outputChannel)

	var ball_x, paddle_x int

	for {
		var x int
		select {
		// Try reading from the intcode computer process output
		case x = <-outputChannel:
		// If it hasn't produced output in a while, assume it's blocked waiting for user input
		case <-time.After(10 * time.Millisecond):
			// Calculate joystick position and provide it as input to the process
			inputChannel <- calcJoystickPosition(ball_x, paddle_x)

			// Consume the output after unblocking the process
			x = <-outputChannel
		}

		y := <-outputChannel

		tile, ok := <-outputChannel
		if !ok {
			break
		}

		if isPaddle := tile == 3; isPaddle {
			paddle_x = x
		} else if isBall := tile == 4; isBall {
			ball_x = x
		}

		if isScoreOutput := x == -1 && y == 0; isScoreOutput {
			fmt.Printf("Score: %d\n", tile)
		}
	}
}
