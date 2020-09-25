package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputFileName string = "input.txt"

func countOrbitalTransfers(orbits map[string]string) int {
	var count int
	pathCount := make(map[string]int)

	cur, ok := orbits["YOU"]
	for ok {
		pathCount[cur] = count
		count++
		cur, ok = orbits[cur]
	}

	count = 0
	cur, ok = orbits["SAN"]
	for ok {
		if found, ok := pathCount[cur]; ok {
			return found + count
		}

		count++

		cur, ok = orbits[cur]
	}

	return 0
}

func countOrbits(orbits map[string]string, orbitter string) int {
	var count int

	orbitter, ok := orbits[orbitter]
	for ok {
		count++
		orbitter, ok = orbits[orbitter]
	}

	return count
}

func readOrbitsFromInputFile() (map[string]string, error) {
	f, err := os.Open(inputFileName)
	if err != nil {
		return nil, fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	orbits := make(map[string]string)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			orbitData := strings.Split(line, ")")
			orbitted := orbitData[0]
			orbitter := orbitData[1]

			orbits[orbitter] = orbitted
		}
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading input file: %w", err)
	}

	return orbits, nil
}

func main() {
	orbits, err := readOrbitsFromInputFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	var totalOrbitsCount int
	for k, _ := range orbits {
		totalOrbitsCount += countOrbits(orbits, k)
	}
	fmt.Println(totalOrbitsCount)

	fmt.Println(countOrbitalTransfers(orbits))
}
