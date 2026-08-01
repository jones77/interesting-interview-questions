package main

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		got := LargestRectangleArea(tt.heights)
		if got != tt.want {
			t.Errorf("LargestRectangleArea(%v) = %v, want %v", tt.heights, got, tt.want)
		}
	}
}
