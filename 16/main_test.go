package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "t1",
			input:  "80871224585914546619083218645595",
			result: "24176176",
		},
		{
			name:   "t2",
			input:  "19617804207202209144916044189917",
			result: "73745418",
		},
		{
			name:   "t3",
			input:  "69317163492948606335995924319873",
			result: "52432133",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, SolvePartOne(tt.input))
		})
	}
}
