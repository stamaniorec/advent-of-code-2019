package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFileName string = "input.txt"

func calculateRequiredFuel(mass int) int {
	return mass/3 - 2
}

func calculateTotalRequiredFuel(mass int) (totalFuel int) {
	fuel := calculateRequiredFuel(mass)

	for fuel > 0 {
		totalFuel += fuel
		fuel = calculateRequiredFuel(fuel)
	}

	return
}

func readMassesFromInputFile() ([]int, error) {
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
			mass, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("Can not convert %s to number: %w", line, err)
			}

			masses = append(masses, mass)
		}
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading input file: %w", err)
	}

	return masses, nil
}

func main() {
	masses, err := readMassesFromInputFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	var totalRequiredFuel int
	for _, mass := range masses {
		// totalRequiredFuel += calculateRequiredFuel(mass)
		totalRequiredFuel += calculateTotalRequiredFuel(mass)
	}

	fmt.Println(totalRequiredFuel)
}
