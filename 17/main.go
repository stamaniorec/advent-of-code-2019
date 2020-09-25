package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	// "strconv"
	"strings"

	"./camera"
	"./graph"
)

func readProgram() (string, error) {
	dat, err := ioutil.ReadFile("input_part2.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func findAlignmentParametersSum(viewport camera.Viewport) int {
	var sum int
	for i := 1; i < viewport.Rows-1; i++ {
		for j := 1; j < viewport.Cols-1; j++ {
			if isScaffold := string(viewport.Grid[i][j]) == "#"; isScaffold {
				if camera.IsScaffoldIntersection(viewport, i, j) {
					sum += findAlignmentParameter(i, j)
				}
			}
		}
	}
	return sum
}

func findAlignmentParameter(i, j int) int {
	return i * j
}

func findIntersections(viewport camera.Viewport) []graph.Node {
	var intersections []graph.Node
	for i := 1; i < viewport.Rows-1; i++ {
		for j := 1; j < viewport.Cols-1; j++ {
			if isScaffold := string(viewport.Grid[i][j]) == "#"; isScaffold {
				if camera.IsScaffoldIntersection(viewport, i, j) {
					intersections = append(intersections, graph.Node{Row: i, Col: j})
				}
			}
		}
	}
	return intersections
}

func printViewportWithEulerPath(viewport camera.Viewport, eulerPath []graph.Node, intersections []graph.Node) {
	var digit int
	for _, u := range eulerPath {
		if u.IsIntersection {
			digit = (digit + 1) % 10
		}

		viewport.Grid[u.Row][u.Col] = digit + '0'
	}

	camera.PrintViewport(viewport)

	for _, v := range eulerPath {
		viewport.Grid[v.Row][v.Col] = '#'
	}
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func part2(program string) ([]int, error) {
	cmd := exec.Command("../09/main", program)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	commands := []string{
		"A,C,A,C,B,B,C,A,C,B\n",
		"L,4,R,8,L,6,L,10\n",
    "L,4,L,4,L,10\n",
		"L,6,R,8,R,10,L,6,L,6\n",
		"n\n",
	}

	for _, c := range commands {
		for _, ch := range c {
			io.WriteString(stdin, fmt.Sprintf("%d\n", int(ch)))
		}
	}
	// io.WriteString(stdin, "n\n")

	for {
		ok := scanner.Scan()
		if !ok {
			break
		}

		outputLine := scanner.Text()
		// outputAsNumber, _ := strconv.Atoi(strings.TrimSpace(outputLine))
		// fmt.Print(string(rune(outputAsNumber)))
		fmt.Println(outputLine)
	}

	return nil, cmd.Wait()
}

func main() {
	program, err := readProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	// viewport, err := camera.ReadViewport(program)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// camera.PrintViewport(viewport)

	// Part 1
	// fmt.Println(findAlignmentParametersSum(viewport))

	// Part 2
	// node := graph.Node{Row: 16, Col: 36}
	// g := graph.BuildGraph(viewport, node)
	// eulerPath := graph.FindEulerPath(g, node)
	// fmt.Println(eulerPath[0])
	// fmt.Println(eulerPath[len(eulerPath)-1])
	// fmt.Println(len(eulerPath))

	// intersections := findIntersections(viewport)
	// printViewportWithEulerPath(viewport, eulerPath, intersections)

	// I dedicated too much time on this problem.
	// So far, I've found the path the robot needs to follow
	// Next steps would be to represent that path in terms of directions
	// e.g. left, move 3, up, move 8, etc.
	// I'm not sure how to tackle the memory constraint complication
	// but the path is short so it should be solvable by brute force

	// graph.FindPath(g, node)

	// I tried this once more, but I just can't think of anything smart enough
	// so that I don't spend a day on it
	// So I solved this manually, which was actually harder than expected
	// because there are many possible paths

	// L4 R8 L6 L10 L6 R8 R10 L6 L6 L4 R8 L6 L10 L6 R8 R10 L6 L6 L4 L4 L10 L4 L4 L10 L6 R8 R10 L6 L6 L4 R8 L6 L10 L6 R8 R10 L6 L6 L4 L4 L10
	// A,C,A,C,B,B,C,A,C,B
	// A="L,4,R,8,L,6,L,10"
  // B="L,4,L,4,L,10"
	// C="L,6,R,8,R,10,L,6,L,6"

	part2(program)
}
