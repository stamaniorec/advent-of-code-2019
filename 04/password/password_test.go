package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password int
		result   bool
	}{
		{
			name:     "t1",
			password: 122345,
			result:   true,
		},
		{
			name:     "t2",
			password: 111123,
			result:   true,
		},
		{
			name:     "t3",
			password: 135679,
			result:   false,
		},
		{
			name:     "t4",
			password: 111111,
			result:   true,
		},
		{
			name:     "t5",
			password: 223450,
			result:   false,
		},
		{
			name:     "t6",
			password: 123789,
			result:   false,
		},
		{
			name:     "t7",
			password: 1233889,
			result:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPassword(tt.password)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestIsValidPasswordPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		password int
		result   bool
	}{
		{
			name:     "t1",
			password: 122345,
			result:   true,
		},
		{
			name:     "t2",
			password: 111123,
			result:   false,
		},
		{
			name:     "t3",
			password: 135679,
			result:   false,
		},
		{
			name:     "t4",
			password: 111111,
			result:   false,
		},
		{
			name:     "t5",
			password: 223450,
			result:   false,
		},
		{
			name:     "t6",
			password: 123789,
			result:   false,
		},
		{
			name:     "t7",
			password: 1233889,
			result:   true,
		},
		{
			name:     "t8",
			password: 112233,
			result:   true,
		},
		{
			name:     "t9",
			password: 123444,
			result:   false,
		},
		{
			name:     "t10",
			password: 111122,
			result:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPasswordPartTwo(tt.password)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestHasAdjacentRepeatingDigits(t *testing.T) {
	tests := []struct {
		name     string
		password int
		result   bool
	}{
		{
			name:     "t1",
			password: 111,
			result:   true,
		},
		{
			name:     "t2",
			password: 112,
			result:   true,
		},
		{
			name:     "t3",
			password: 122,
			result:   true,
		},
		{
			name:     "t4",
			password: 159,
			result:   false,
		},
		{
			name:     "t5",
			password: 122345,
			result:   true,
		},
		{
			name:     "t6",
			password: 135679,
			result:   false,
		},
		{
			name:     "t7",
			password: 221,
			result:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasAdjacentRepeatingDigits(tt.password)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestHasNonDecreasingDigits(t *testing.T) {
	tests := []struct {
		name     string
		password int
		result   bool
	}{
		{
			name:     "t1",
			password: 111,
			result:   true,
		},
		{
			name:     "t2",
			password: 112,
			result:   true,
		},
		{
			name:     "t3",
			password: 122,
			result:   true,
		},
		{
			name:     "t4",
			password: 159,
			result:   true,
		},
		{
			name:     "t5",
			password: 111123,
			result:   true,
		},
		{
			name:     "t6",
			password: 135679,
			result:   true,
		},
		{
			name:     "t7",
			password: 221,
			result:   false,
		},
		{
			name:     "t8",
			password: 223450,
			result:   false,
		},
		{
			name:     "t9",
			password: 224399,
			result:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasNonDecreasingDigits(tt.password)
			assert.Equal(t, tt.result, result)
		})
	}
}
