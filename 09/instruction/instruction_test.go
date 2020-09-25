package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParameterMode(t *testing.T) {
	tests := []struct {
		name           string
		instruction    int
		parameterIndex int
		parameterMode  ParameterMode
	}{
		{
			name:           "t1",
			instruction:    1,
			parameterIndex: 0,
			parameterMode:  PositionalMode,
		},
		{
			name:           "t2",
			instruction:    1,
			parameterIndex: 1,
			parameterMode:  PositionalMode,
		},
		{
			name:           "t3",
			instruction:    101,
			parameterIndex: 0,
			parameterMode:  ImmediateMode,
		},
		{
			name:           "t4",
			instruction:    101,
			parameterIndex: 1,
			parameterMode:  PositionalMode,
		},
		{
			name:           "t5",
			instruction:    1101,
			parameterIndex: 1,
			parameterMode:  ImmediateMode,
		},
		{
			name:           "t6",
			instruction:    1101,
			parameterIndex: 2,
			parameterMode:  PositionalMode,
		},
		{
			name:           "t7",
			instruction:    204,
			parameterIndex: 0,
			parameterMode:  RelativeMode,
		},
		{
			name:           "t8",
			instruction:    2204,
			parameterIndex: 1,
			parameterMode:  RelativeMode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameterMode := GetParameterMode(tt.instruction, tt.parameterIndex)
			assert.Equal(t, tt.parameterMode, parameterMode)
		})
	}
}
