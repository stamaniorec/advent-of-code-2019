package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	v "./vector"
)

const NumMoons = 4

func parseMoonPosition(inputLine string) v.Vector {
	regex := regexp.MustCompile(`<x=(?P<X>-?\d+), y=(?P<Y>-?\d+), z=(?P<Z>-?\d+)>`)
	match := regex.FindStringSubmatch(inputLine)

	position := make(map[string]int)
	for i, coord := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			value, _ := strconv.Atoi(match[i])
			position[coord] = value
		}
	}

	return v.Vector{
		X: position["X"],
		Y: position["Y"],
		Z: position["Z"],
	}
}

func readMoons() ([NumMoons]v.Vector, error) {
	var moons [NumMoons]v.Vector

	f, err := os.Open("input.txt")
	if err != nil {
		return moons, fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for i := 0; i < NumMoons; i++ {
		ok := sc.Scan()
		if !ok {
			return moons, errors.New("Error while reading file")
		}

		line := sc.Text()
		if line != "" {
			moons[i] = parseMoonPosition(strings.TrimSpace(line))
		}
	}
	if err := sc.Err(); err != nil {
		return moons, fmt.Errorf("Error while reading input file: %w", err)
	}

	return moons, nil
}

// https://github.com/prscoelho/aoc2019/blob/master/src/aoc12/mod.rs
//https://www.reddit.com/r/adventofcode/comments/e9jxh2/help_2019_day_12_part_2_what_am_i_not_seeing/

func calculateGravityOnAxis(current, other int) int {
	switch {
	case current > other:
		return -1
	case current < other:
		return 1
	default:
		return 0
	}
}

func calculateGravity(a, b v.Vector) v.Vector {
	return v.Vector{
		X: calculateGravityOnAxis(a.X, b.X),
		Y: calculateGravityOnAxis(a.Y, b.Y),
		Z: calculateGravityOnAxis(a.Z, b.Z),
	}
}

func calculateTotalGravity(moon v.Vector, moons [NumMoons]v.Vector) v.Vector {
	var totalGravity v.Vector
	for _, m := range moons {
		totalGravity = v.AddVectors(totalGravity, calculateGravity(moon, m))
	}
	return totalGravity
}

func calculateTotalGravityOnAxis(value int, allValues [NumMoons]int) int {
	var totalGravity int
	for _, x := range allValues {
		totalGravity += calculateGravityOnAxis(value, x)
	}
	return totalGravity
}

func calculateVelocities(moons, currentVelocities [NumMoons]v.Vector) [NumMoons]v.Vector {
	var velocities [NumMoons]v.Vector
	for i, m := range moons {
		totalGravity := calculateTotalGravity(m, moons)
		velocities[i] = v.AddVectors(currentVelocities[i], totalGravity)
	}
	return velocities
}

func calculateVelocitiesOnAxis(values, currentVelocities [NumMoons]int) [NumMoons]int {
	for i, m := range values {
		totalGravity := calculateTotalGravityOnAxis(m, values)
		currentVelocities[i] += totalGravity
	}
	return currentVelocities
}

func getEnergy(vec v.Vector) int {
	av := v.GetAbsoluteVector(vec)
	return av.X + av.Y + av.Z
}

func calculateTotalEnergy(moons, velocities [NumMoons]v.Vector) int {
	var totalEnergy int
	for i := range moons {
		potentialEnergy := getEnergy(moons[i])
		kineticEnergy := getEnergy(velocities[i])
		totalEnergy += potentialEnergy * kineticEnergy
	}
	return totalEnergy
}

func lcm(nums ...int) int {
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		ans = (nums[i] * ans) / gcd(nums[i], ans)
	}

	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func timesUntilRepeatedState(moons, velocities [NumMoons]v.Vector) int {
	xs, dxs := [NumMoons]int{}, [NumMoons]int{}
	ys, dys := [NumMoons]int{}, [NumMoons]int{}
	zs, dzs := [NumMoons]int{}, [NumMoons]int{}
	for i := range moons {
		xs[i] = moons[i].X
		ys[i] = moons[i].Y
		zs[i] = moons[i].Z

		dxs[i] = velocities[i].X
		dys[i] = velocities[i].Y
		dzs[i] = velocities[i].Z
	}

	xSteps := calculateStepsBeforeLoop(xs, dxs)
	ySteps := calculateStepsBeforeLoop(ys, dys)
	zSteps := calculateStepsBeforeLoop(zs, dzs)

	return lcm(xSteps, ySteps, zSteps)
}

func calculateStateHash(moonPositionsByAxis, velocitiesByAxis [NumMoons]int) string {
	var hash string
	for i := range moonPositionsByAxis {
		hash += fmt.Sprintf("%d-%d|", moonPositionsByAxis[i], velocitiesByAxis[i])
	}
	return hash
}

func calculateStepsBeforeLoop(moonPositionsByAxis, velocitiesByAxis [NumMoons]int) int {
	var steps int

	states := map[string]bool{}
	for {
		stateHash := calculateStateHash(moonPositionsByAxis, velocitiesByAxis)
		if states[stateHash] {
			break
		}

		states[stateHash] = true

		velocitiesByAxis = calculateVelocitiesOnAxis(moonPositionsByAxis, velocitiesByAxis)

		for i := range moonPositionsByAxis {
			moonPositionsByAxis[i] += velocitiesByAxis[i]
		}

		steps++
	}

	return steps
}

func main() {
	moons, err := readMoons()
	if err != nil {
		fmt.Println(err)
		return
	}

	var velocities [NumMoons]v.Vector

	for step := 0; step < 1000; step++ {
		velocities = calculateVelocities(moons, velocities)

		for i := range moons {
			moons[i] = v.AddVectors(moons[i], velocities[i])
			// fmt.Printf("pos - %d %d %d\n", moons[i].X, moons[i].Y, moons[i].Z)
		}
		// fmt.Println()
	}

	fmt.Printf("Total energy: %d\n", calculateTotalEnergy(moons, velocities))

	steps := timesUntilRepeatedState(moons, velocities)
	fmt.Println(steps)
}
