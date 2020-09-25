package main

import (
	"bufio"
	"fmt"
	"os"

	"./intersection"
)

const inputFileName string = "input.txt"

func readPathsFromFile(fileName string) (wire1Path, wire2Path string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", "", fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	sc.Scan()
	wire1Path = sc.Text()

	if err := sc.Err(); err != nil {
		return "", "", fmt.Errorf("Error while reading first line of file: %w", err)
	}

	sc.Scan()
	wire2Path = sc.Text()

	if err := sc.Err(); err != nil {
		return "", "", fmt.Errorf("Error while reading second line of file: %w", err)
	}

	return
}

func main() {
	wire1Path, wire2Path, err := readPathsFromFile(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	intersections, err := intersection.FindIntersections(wire1Path, wire2Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Manhattan Distance: %d\n",
		intersection.MinManhattanDistanceToIntersection(intersections))

	fmt.Printf("Fewest steps from central point to intersection: %d\n",
		intersection.FewestStepsFromCentralPointIntersection(intersections))
}
