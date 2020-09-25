package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"./droid"
)

func readProgram() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

type Vector struct {
	X, Y int
}

type QueueItem struct {
	Position Vector
	Distance int
}

var directions = []int{1, 2, 3, 4}
var directionVectors = []Vector{
	Vector{0, -1},
	Vector{0, 1},
	Vector{-1, 0},
	Vector{1, 0},
}

var distances = []int{}

func main() {
	program, err := readProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputChannel := make(chan int)
	outputChannel := make(chan int)

	go droid.Run(program, inputChannel, outputChannel)

	startingPosition := Vector{
		X: 0,
		Y: 0,
	}

	visited := make(map[Vector]bool)

	dfs(startingPosition, 0, visited, inputChannel, outputChannel, 0)

	fmt.Println(distances)

	oxygenSource := Vector{
		X: -12,
		Y: 18,
	}

	minutes := countMinutes(oxygenSource, visited)
	fmt.Println(minutes)
}

func countMinutes(oxygenSource Vector, visited map[Vector]bool) int {
	queue := []QueueItem{
		QueueItem{
			Position: oxygenSource,
			Distance: 0,
		},
	}

	var minutes int
	var lastDistance int

	for len(queue) > 0 {
		queueItem := queue[0]
		queue = queue[1:len(queue)]

		if queueItem.Distance > lastDistance {
			minutes++
			lastDistance = queueItem.Distance
		}

		for i, _ := range directions {
			position := queueItem.Position
			nextPosition := Vector{
				X: position.X + directionVectors[i].X,
				Y: position.Y + directionVectors[i].Y,
			}

			if _, found := visited[nextPosition]; !found {
				continue
			}

			delete(visited, nextPosition)

			queue = append(queue, QueueItem{
				Position: nextPosition,
				Distance: queueItem.Distance + 1,
			})
		}
	}

	return minutes
}

func getReverseDirection(direction int) int {
	if direction == 1 {
		return 2
	} else if direction == 2 {
		return 1
	} else if direction == 3 {
		return 4
	} else if direction == 4 {
		return 3
	}

	return 0
}

func dfs(position Vector, incomingDir int, visited map[Vector]bool, inputChannel chan int, outputChannel chan int, distance int) {
	visited[position] = true

	for i, d := range directions {
		nextPosition := Vector{
			X: position.X + directionVectors[i].X,
			Y: position.Y + directionVectors[i].Y,
		}

		if _, found := visited[nextPosition]; found {
			continue
		}

		inputChannel <- d
		reply := <-outputChannel

		if isDestination := reply == 2; isDestination {
			fmt.Println("found destination")
			fmt.Println(nextPosition)
			distances = append(distances, distance+1)

			reverseDirection := getReverseDirection(d)
			inputChannel <- reverseDirection
			<-outputChannel
			continue
		}

		if isWall := reply == 0; isWall {
			continue
		}

		dfs(nextPosition, d, visited, inputChannel, outputChannel, distance+1)
	}

	reverseDirection := getReverseDirection(incomingDir)

	inputChannel <- reverseDirection
	<-outputChannel
	// }
}
