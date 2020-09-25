package main

import (
	"fmt"

	"./password"
)

func CountValidPasswords(rangeStart, rangeEnd int) int {
	var count int

	for i := rangeStart; i <= rangeEnd; i++ {
		if password.IsValidPasswordPartTwo(i) {
			count++
		}
	}

	return count
}

func main() {
	rangeStart := 387638
	rangeEnd := 919123

	fmt.Println(CountValidPasswords(rangeStart, rangeEnd))
}
