package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOrbits(t *testing.T) {
	orbits := map[string]string{
		"D": "C",
		"E": "D",
		"G": "B",
		"H": "G",
		"I": "D",
		"B": "COM",
		"F": "E",
		"J": "E",
		"K": "J",
		"L": "K",
		"C": "B",
	}

	tests := []struct {
		name        string
		orbits      map[string]string
		orbitter    string
		orbitsCount int
	}{
		{
			name:        "t1",
			orbits:      orbits,
			orbitter:    "COM",
			orbitsCount: 0,
		},
		{
			name:        "t2",
			orbits:      orbits,
			orbitter:    "B",
			orbitsCount: 1,
		},
		{
			name:        "t3",
			orbits:      orbits,
			orbitter:    "D",
			orbitsCount: 3,
		},
		{
			name:        "t4",
			orbits:      orbits,
			orbitter:    "L",
			orbitsCount: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.orbitsCount, countOrbits(tt.orbits, tt.orbitter))
		})
	}
}

func TestCountOrbitalTransfers(t *testing.T) {
	orbits := map[string]string{
		"B":   "COM",
		"C":   "B",
		"D":   "C",
		"E":   "D",
		"F":   "E",
		"G":   "B",
		"H":   "G",
		"I":   "D",
		"J":   "E",
		"K":   "J",
		"L":   "K",
		"YOU": "K",
		"SAN": "I",
	}

	assert.Equal(t, 4, countOrbitalTransfers(orbits))
}
