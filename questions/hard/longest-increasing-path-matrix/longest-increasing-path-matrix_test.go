package longestincreasingpathmatrix_test

import (
	"github.com/jones/interesting-interview-questions/questions/hard/longest-increasing-path-matrix"
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
		got := longestincreasingpathmatrix.LongestIncreasingPath(tt.matrix)
		if got != tt.want {
			t.Errorf("longestincreasingpathmatrix.LongestIncreasingPath(%v) = %d, want %d", tt.matrix, got, tt.want)
		}
	}
}
