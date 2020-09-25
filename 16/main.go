package main

import (
	"strconv"
	"strings"
)

func splitToDigitsList(input string) []int {
	list := make([]int, len(input))

	digits := strings.Split(input, "")
	for i, d := range digits {
		parsedDigit, _ := strconv.Atoi(d)
		list[i] = parsedDigit
	}

	return list
}

var multiplierBasePattern = []int{0, 1, 0, -1}

func calculateElement(index int, inputList []int) int {
	var outputElement int

	var sameCount int = index
	var multiplierIndex int
	for i := 0; i < len(inputList); i++ {
		if sameCount == 0 {
			multiplierIndex++
			multiplierIndex = multiplierIndex % len(multiplierBasePattern)
			sameCount = index + 1
		}

		outputElement += inputList[i] * multiplierBasePattern[multiplierIndex]

		sameCount--
	}
	outputElement = outputElement % 10
	if outputElement < 0 {
		outputElement = -outputElement
	}

	return outputElement
}

func FindMessage(input string) string {
	inputList := splitToDigitsList(input)

	var result string

	for i := 0; i < 100; i++ {
		outputList := make([]int, len(inputList))

		for j := 0; j < len(outputList); j++ {
			outputList[j] = calculateElement(j, inputList)
		}

		var stringed []string
		for _, v := range outputList {
			stringed = append(stringed, strconv.Itoa(v))
		}
		result = strings.Join(stringed[:], "")

		copy(inputList, outputList)
	}

	return result
}

func SolvePartOne(input string) string {
	msg := FindMessage(input)
	return msg[:8]
}

func SolvePartTwo(input string) string {
	// partOneInput := "03036732577212944063491565474664"
	messageOffset, _ := strconv.Atoi(input[:7])

	input = strings.Repeat(input, 10000)

	msg := FindMessage(input)
	return msg[messageOffset : messageOffset+8]
}

func main() {
	// fmt.Println(SolvePartTwo("02935109699940807407585447034323"))
}
