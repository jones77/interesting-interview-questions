package main

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{
			[][]byte{{'0'}},
			0,
		},
	}
	for _, tt := range tests {
		got := NumIslands(tt.grid)
		if got != tt.want {
			t.Errorf("NumIslands(%v) = %d, want %d", tt.grid, got, tt.want)
		}
	}
}
