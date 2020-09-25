package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"./droid"
)

/*
Take festive hat
Take whirled peas
Take pointer
Take dark matter
Take mutex
*/

func readProgram() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return string(dat), nil
}

func main() {
	program, err := readProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	commandChannel := make(chan string)

	go droid.Run(program, commandChannel)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLine := scanner.Text()

		commandChannel <- inputLine
	}
}
