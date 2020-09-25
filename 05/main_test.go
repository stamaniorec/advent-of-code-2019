package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteIntcodeProgram(t *testing.T) {
	tests := []struct {
		name    string
		program []int
		result  int
	}{
		{
			name:    "t1",
			program: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			result:  3500,
		},
		{
			name:    "t2",
			program: []int{1, 0, 0, 0, 99},
			result:  2,
		},
		{
			name:    "t3",
			program: []int{1, 1, 2, 0, 99},
			result:  3,
		},
		{
			name:    "t4",
			program: []int{2, 1, 2, 0, 99},
			result:  2,
		},
		{
			name:    "t5",
			program: []int{2, 4, 4, 5, 99, 0},
			result:  2,
		},
		{
			name:    "t6",
			program: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			result:  30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExecuteIntcodeProgram(tt.program)
			assert.Equal(t, tt.result, tt.program[0])
		})
	}
}
