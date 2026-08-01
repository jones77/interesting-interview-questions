package main

import (
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
		{[][]int{{1}}, 1},
	}
	for _, tt := range tests {
		got := LongestIncreasingPath(tt.matrix)
		if got != tt.want {
			t.Errorf("LongestIncreasingPath(%v) = %d, want %d", tt.matrix, got, tt.want)
		}
	}
}
