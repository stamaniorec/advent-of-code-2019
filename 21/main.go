package main

import (
	"fmt"
	"io/ioutil"

	"./droid"
)

var part1 = []string{
	"NOT A T",
	"OR T J",
	"NOT B T",
	"OR T J",
	"NOT C T",
	"OR T J",
	"AND D J",
	"WALK",
}

var part2 = []string{
	"NOT A T",
	"OR T J",
	"NOT B T",
	"OR T J",
	"NOT C T",
	"OR T J",
	"AND D J",

	"NOT J T",
	"OR H T",
	"OR E T",
	"AND T J",
	"RUN",
}

func readInput() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return string(dat), nil
}

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := droid.Run(input, part2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
