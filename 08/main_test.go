package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLayerWithFewestZeroes(t *testing.T) {
	tests := []struct {
		name                   string
		imageData              string
		width                  int
		height                 int
		fewestZeroesLayerIndex int
	}{
		{
			name:                   "t1",
			imageData:              "123456",
			width:                  3,
			height:                 2,
			fewestZeroesLayerIndex: 0,
		},
		{
			name:                   "t2",
			imageData:              "123456789012",
			width:                  3,
			height:                 2,
			fewestZeroesLayerIndex: 0,
		},
		{
			name:                   "t3",
			imageData:              "120456789112",
			width:                  3,
			height:                 2,
			fewestZeroesLayerIndex: 1,
		},
		{
			name:                   "t4",
			imageData:              "12341234",
			width:                  2,
			height:                 2,
			fewestZeroesLayerIndex: 0,
		},
		{
			name:                   "t4",
			imageData:              "000010041234",
			width:                  2,
			height:                 2,
			fewestZeroesLayerIndex: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layerIndex := GetLayerWithFewestZeroes(tt.imageData, tt.width, tt.height)
			assert.Equal(t, tt.fewestZeroesLayerIndex, layerIndex)
		})
	}
}
