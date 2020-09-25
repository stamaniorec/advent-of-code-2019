package solution

import (
	"math"
	"unicode"

	g "../grid"
)

// KeyResult represents a key with Label, located at Position
// which is NumMoves moves away and is reachable by the robot at StartingRobotPos
type KeyResult struct {
	Position         g.Position
	Label            rune
	NumMoves         int
	StartingRobotPos g.Position
}

// GetKeyCandidatesDFS finds the keys that can be picked from the current position
// along with some
func GetKeyCandidatesBFS(grid [][]rune, robotPositions []g.Position) []KeyResult {
	type queueItem struct {
		pos              g.Position
		prevPos          g.Position
		StartingRobotPos g.Position
		moves            int
	}

	var q []queueItem
	for _, rp := range robotPositions {
		q = append(q, queueItem{
			pos:              rp,
			prevPos:          g.Position{Row: -1, Col: -1},
			StartingRobotPos: rp,
			moves:            0,
		})
	}

	rowsCount := len(grid)
	colsCount := len(grid[0])

	var keys []KeyResult

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		nextPositions := []g.Position{
			{Row: item.pos.Row - 1, Col: item.pos.Col},
			{Row: item.pos.Row, Col: item.pos.Col - 1},
			{Row: item.pos.Row, Col: item.pos.Col + 1},
			{Row: item.pos.Row + 1, Col: item.pos.Col},
		}

		for _, nextPos := range nextPositions {
			if nextPos == item.prevPos {
				continue
			}

			isWithinBounds :=
				nextPos.Row >= 0 && nextPos.Col >= 0 &&
					nextPos.Row < rowsCount && nextPos.Col < colsCount
			if !isWithinBounds {
				continue
			}

			if isKey := unicode.IsLower(grid[nextPos.Row][nextPos.Col]); isKey {
				keys = append(keys, KeyResult{
					Position:         nextPos,
					Label:            grid[nextPos.Row][nextPos.Col],
					NumMoves:         item.moves + 1,
					StartingRobotPos: item.StartingRobotPos,
				})

				continue
			}

			if isFree := grid[nextPos.Row][nextPos.Col] == '.'; isFree {
				q = append(q, queueItem{
					pos:              nextPos,
					prevPos:          item.pos,
					moves:            item.moves + 1,
					StartingRobotPos: item.StartingRobotPos,
				})
			}
		}
	}

	return keys
}

// GetKeyCandidatesDFS finds the keys that can be picked from the current position
// I initially wrote the DFS, but for bigger inputs I started getting stack overflow
// so I wrote a BFS instead
func GetKeyCandidatesDFS(grid [][]rune,
	pos g.Position, prevPos g.Position, moves int) []KeyResult {

	var keys []KeyResult

	isKey := unicode.IsLower(grid[pos.Row][pos.Col])
	if isKey {
		keys = append(keys, KeyResult{
			Position: pos,
			Label:    grid[pos.Row][pos.Col],
			NumMoves: moves,
		})
	}

	next := []g.Position{
		{Row: pos.Row - 1, Col: pos.Col},
		{Row: pos.Row, Col: pos.Col - 1},
		{Row: pos.Row, Col: pos.Col + 1},
		{Row: pos.Row + 1, Col: pos.Col},
	}

	allRows := len(grid)
	allCols := len(grid[0])

	for _, x := range next {
		isWithinBounds :=
			x.Row >= 0 && x.Col >= 0 &&
				x.Row < allRows && x.Col < allCols

		if isWithinBounds {
			isFree := grid[x.Row][x.Col] == '.'
			isKey := grid[x.Row][x.Col] >= 'a' && grid[x.Row][x.Col] <= 'z'

			if (isFree || isKey) && x != prevPos {
				keys = append(keys, GetKeyCandidatesDFS(grid, x, pos, moves+1)...)
			}
		}
	}

	return keys
}

// FindMinSteps solves the problem with a memoized backtrack
func FindMinSteps(grid [][]rune,
	robotPositions []g.Position, remainingKeys int, doors map[rune]g.Position, dp map[int64]int) int {

	hash := g.GetGridHash(grid)

	// keys := getKeyCandidatesDFS(grid, pos, pos, 0)
	keysCandidates := GetKeyCandidatesBFS(grid, robotPositions)

	// Base case
	if aKeyCanBePicked := len(keysCandidates) == 0; aKeyCanBePicked {
		if remainingKeys == 0 {
			dp[hash] = 0
			return 0
		} else {
			panic("No keys available to be picked, but there are still keys remaining!")
		}
	}

	minSteps := math.MaxInt32
	for _, key := range keysCandidates {
		// Obtain key by moving robot i
		for i := range robotPositions {
			if robotPositions[i] == key.StartingRobotPos {
				robotPositions[i] = key.Position
			}
		}
		g.ObtainKey(grid, key.Position, key.StartingRobotPos, doors)

		// Backtrack
		var moves int
		hash := g.GetGridHash(grid)
		if _, ok := dp[hash]; ok {
			moves = dp[hash]
		} else {
			moves = FindMinSteps(grid, robotPositions, remainingKeys-1, doors, dp)
		}

		if key.NumMoves+moves < minSteps {
			minSteps = key.NumMoves + moves
		}

		// Put back the key and move robot i back to where it was
		for i := range robotPositions {
			if robotPositions[i] == key.Position {
				robotPositions[i] = key.StartingRobotPos
			}
		}
		g.PutBackKey(grid, key.Position, key.StartingRobotPos, key.Label, doors)
	}

	dp[hash] = minSteps

	return minSteps
}
