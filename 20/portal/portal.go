package portal

import (
	"unicode"
	"regexp"
	"strings"

	d "../direction"
	p "../maze/position"
	m "../maze"
	"../util"
)

type ReadConfig struct {
	TraverseDir d.Direction
	ReadDir     d.Direction
	Reverse     bool
}

type Portal struct {
	Label    string
	Position p.Position
}

func readPortalLabel(maze m.Maze, pos p.Position, dirVect p.Position, reverse bool) string {
	var sb strings.Builder
	for {
		pos.Row += dirVect.Row
		pos.Col += dirVect.Col

		if unicode.IsControl(maze.Grid[pos.Row][pos.Col]) {
			break
		}
		if unicode.IsSpace(maze.Grid[pos.Row][pos.Col]) {
			break
		}

		sb.WriteRune(maze.Grid[pos.Row][pos.Col])
	}

	str := sb.String()
	if reverse {
		str = util.ReverseString(str)
	}

	return str
}

func FindPortals(maze m.Maze, pos p.Position, conf ReadConfig) []Portal {
	var portals []Portal
	dirVect := d.DirectionVectors[conf.TraverseDir]

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	for pos.Row < len(maze.Grid) && pos.Col < len(maze.Grid[pos.Row]) {
		if maze.Grid[pos.Row][pos.Col] == '.' {
			label := readPortalLabel(maze, pos, d.DirectionVectors[conf.ReadDir], conf.Reverse)

			if isAlpha(label) {
				portals = append(portals, Portal{
					Label:    label,
					Position: pos,
				})
			}
		}

		pos.Row += dirVect.Row
		pos.Col += dirVect.Col
	}

	return portals
}

func FindOuterRingPortals(maze m.Maze) []Portal {
	reads := []struct{
		StartingPos p.Position
		ReadSettings ReadConfig
	}{
		{
			StartingPos: maze.TopLeft,
			ReadSettings: ReadConfig{
				TraverseDir: d.Down,
				ReadDir:     d.Left,
				Reverse:     true,
			},
		},

		{
			StartingPos: maze.TopLeft,
			ReadSettings: ReadConfig{
				TraverseDir: d.Right,
				ReadDir:     d.Up,
				Reverse:     true,
			},
		},

		{
			StartingPos: p.Position {
				Row: maze.TopLeft.Row + maze.Height-1,
				Col: maze.TopLeft.Col,
			},
			ReadSettings: ReadConfig{
				TraverseDir: d.Right,
				ReadDir:     d.Down,
				Reverse:     false,
			},
		},

		{
			StartingPos: p.Position {
				Row: maze.TopLeft.Row,
				Col: maze.TopLeft.Col + maze.Width-1,
			},
			ReadSettings: ReadConfig{
				TraverseDir: d.Down,
				ReadDir:     d.Right,
				Reverse:     false,
			},
		},
	}

	var portals []Portal
	for _, r := range reads {
		ps := FindPortals(maze, r.StartingPos, r.ReadSettings)
		portals = append(portals, ps...)
	}
	return portals
}

func FindInnerRingPortals(maze m.Maze) []Portal {
	reads := []struct{
		StartingPos p.Position
		ReadSettings ReadConfig
	}{
		{
			StartingPos: maze.InnerTopLeft,
			ReadSettings: ReadConfig{
				TraverseDir: d.Right,
				ReadDir:     d.Down,
				Reverse:     false,
			},
		},

		{
			StartingPos: maze.InnerTopLeft,
			ReadSettings: ReadConfig{
				TraverseDir: d.Down,
				ReadDir:     d.Right,
				Reverse:     false,
			},
		},

		{
			StartingPos: p.Position {
				Row: maze.InnerTopLeft.Row + maze.InnerHeight+1,
				Col: maze.InnerTopLeft.Col,
			},
			ReadSettings: ReadConfig{
				TraverseDir: d.Right,
				ReadDir:     d.Up,
				Reverse:     true,
			},
		},

		{
			StartingPos: p.Position {
				Row: maze.InnerTopLeft.Row,
				Col: maze.InnerTopLeft.Col + maze.InnerWidth+1,
			},
			ReadSettings: ReadConfig{
				TraverseDir: d.Down,
				ReadDir:     d.Left,
				Reverse:     true,
			},
		},
	}

	var innerPortals []Portal
	for _, r := range reads {
		ps := FindPortals(maze, r.StartingPos, r.ReadSettings)
		innerPortals = append(innerPortals, ps...)
	}
	return innerPortals
}

func GetPortal(portals []Portal, label string) Portal {
	for _, x := range portals {
		if x.Label == label {
			return x
		}
	}
	return Portal{}
}

func GetPortalAt(portals []Portal, pos p.Position) Portal {
	for _, x := range portals {
		if x.Position == pos {
			return x
		}
	}
	return Portal{}
}

func FindOppositePortal(portals []Portal, portal Portal) Portal {
	for _, x := range portals {
		if x.Label == portal.Label && x.Position != portal.Position {
			return x
		}
	}
	return Portal{}
}

