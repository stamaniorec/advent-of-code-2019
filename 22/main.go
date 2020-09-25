package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func dealIntoNewStack(deck []int) []int {
	n := len(deck)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[n-i-1] = deck[i]
	}
	return result
}

func cutK(deck []int, k int) []int {
	n := len(deck)
	result := make([]int, 0, n)
	if k >= 0 {
		result = append(result, deck[k:n]...)
		result = append(result, deck[0:k]...)
	} else {
		result = append(result, deck[n+k:n]...)
		result = append(result, deck[0:n+k]...)
	}
	return result
}

func dealWithIncrementK(deck []int, k int) []int {
	result := make([]int, len(deck))
	var i int
	for len(deck) > 0 {
		if result[i] != 0 {
			panic("error")
		}
		result[i] = deck[0]
		i = (i + k) % len(result)
		deck = deck[1:len(deck)]
	}
	return result
}

func readInput() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return string(dat), nil
}

func part1(input string) {
	const deckSize = 10007

	deck := make([]int, 0, deckSize)
	for i := 0; i < deckSize; i++ {
		deck = append(deck, i)
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		op := scanner.Text()
		if strings.HasPrefix(op, "deal with increment") {
			op = strings.ReplaceAll(op, "deal with increment", "")
			op = strings.TrimSpace(op)
			k, _  := strconv.Atoi(op)
			deck = dealWithIncrementK(deck, k)
		} else if strings.HasPrefix(op, "cut") {
			op = strings.ReplaceAll(op, "cut", "")
			op = strings.TrimSpace(op)
			k, _  := strconv.Atoi(op)
			deck = cutK(deck, k)
		} else if op == "deal into new stack" {
			deck = dealIntoNewStack(deck)
		}
	}

	for i, c := range deck {
		if c == 2019 {
			fmt.Println(i)
		}
	}
}

// I tried several approaches with Part 2
// but this is ridiculous
// I can't spend more time on this

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(input)
	part2_2(input)
}
