package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	v "./vector"
)

const inputFileName string = "input.txt"

// Reads the asteroid positions from the input file and returns them as vectors relative to (0,0)
func readAsteroidPositions() ([]v.Vector, error) {
	f, err := os.Open(inputFileName)
	if err != nil {
		return nil, fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	var (
		positions []v.Vector
		y         int
	)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		for x, char := range line {
			if char == '#' {
				positions = append(positions, v.GetPositionVector(x, y))
			}
		}

		y++
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading input file: %w", err)
	}

	return positions, nil
}

func getVectorsRelativeTo(base v.Vector, asteroids []v.Vector) []v.Vector {
	var rvs []v.Vector

	for _, a := range asteroids {
		v := v.GetVectorRelativeTo(base, a)

		if isSelf := v.X == 0 && v.Y == 0; !isSelf {
			rvs = append(rvs, v)
		}
	}

	return rvs
}

func countVisibleAsteroids(base v.Vector, asteroids []v.Vector) int {
	rvs := getVectorsRelativeTo(base, asteroids)
	vis := make(map[v.Vector]struct{})

	for _, rv := range rvs {
		rv = v.GetScaledDownVector(rv)
		if _, ok := vis[rv]; !ok {
			vis[rv] = struct{}{}
		}
	}

	return len(vis)
}

func convertAngleToClockwiseAngle(angleInDegrees float64) float64 {
	switch {
	case angleInDegrees >= 0 && angleInDegrees <= 90:
		return 90 - angleInDegrees
	case angleInDegrees > 90 && angleInDegrees <= 180:
		return 360 - (angleInDegrees - 90)
	case angleInDegrees > 180 && angleInDegrees <= 270:
		return 450 - angleInDegrees
	case angleInDegrees > 270 && angleInDegrees <= 360:
		return 180 - (angleInDegrees - 270)
	default:
		return 0 // error
	}
}

func sortByAngleThenByDistance(origin, a, b v.Vector) bool {
	angleA := v.GetAngle(origin, a)
	adjustedAngleA := convertAngleToClockwiseAngle(v.ToDegrees(angleA))

	angleB := v.GetAngle(origin, b)
	adjustedAngleB := convertAngleToClockwiseAngle(v.ToDegrees(angleB))

	if adjustedAngleA == adjustedAngleB {
		d1 := v.GetMagnitude(v.GetVectorRelativeTo(origin, a))
		d2 := v.GetMagnitude(v.GetVectorRelativeTo(origin, b))
		return d1 <= d2
	} else {
		return adjustedAngleA <= adjustedAngleB
	}
}

func skipRemainingAsteroidsAtSameAngle(base v.Vector, asteroids []v.Vector, i int) int {
	asteroidAngle := v.GetAngle(base, asteroids[i])

	for {
		if i+1 >= len(asteroids) {
			return 0
		}

		nextAsteroid := asteroids[i+1]
		nextAsteroidAngle := v.GetAngle(base, nextAsteroid)

		i++

		if asteroidAngle != nextAsteroidAngle {
			break
		}
	}

	return i
}

func getNthVaporizedAsteroid(base v.Vector, asteroids []v.Vector, n int) v.Vector {
	sort.Slice(asteroids, func(i, j int) bool {
		return sortByAngleThenByDistance(base, asteroids[i], asteroids[j])
	})

	var vapCount, i int
	numAsteroids := len(asteroids)
	numAsteroidsForVaporization := numAsteroids - 1 // except base
	isVaporized := make([]bool, numAsteroids)

	for vapCount < numAsteroidsForVaporization {
		asteroid := asteroids[i]

		if shouldSkip := asteroid == base || isVaporized[i]; shouldSkip {
			i++
			continue
		}

		isVaporized[i] = true
		vapCount++

		if vapCount == n {
			return asteroid
		}

		i = skipRemainingAsteroidsAtSameAngle(base, asteroids, i)
	}

	return v.Vector{}
}

func main() {
	var base v.Vector
	maxVisibleAsteroids := math.MinInt32

	asteroids, _ := readAsteroidPositions()
	for _, a := range asteroids {
		c := countVisibleAsteroids(a, asteroids)

		if c > maxVisibleAsteroids {
			maxVisibleAsteroids = c
			base = a
		}
	}

	fmt.Printf("%f,%f - %d\n", base.Y, base.X, maxVisibleAsteroids)

	nthVaporizedAsteroid := getNthVaporizedAsteroid(base, asteroids, 200)
	result := nthVaporizedAsteroid.X*100 + nthVaporizedAsteroid.Y
	fmt.Printf("%f,%f\n", nthVaporizedAsteroid.X, nthVaporizedAsteroid.Y)
	fmt.Printf("%f\n", result)
}
