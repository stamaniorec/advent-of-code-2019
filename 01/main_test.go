package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuel(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.mass, tt.fuel)

		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tt.fuel, calculateRequiredFuel(tt.mass))
		})
	}
}

func TestCalculateTotalRequiredFuel(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.mass, tt.fuel)

		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tt.fuel, calculateTotalRequiredFuel(tt.mass))
		})
	}
}
