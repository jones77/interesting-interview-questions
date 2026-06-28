package minimumpathsum_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/minimum-path-sum"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{[][]int{{5}}, 5},
	}
	for _, tt := range tests {
		got := minimumpathsum.MinPathSum(tt.grid)
		if got != tt.want {
			t.Errorf("minimumpathsum.MinPathSum(%v) = %d, want %d", tt.grid, got, tt.want)
		}
	}
}
