package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, tt := range tests {
		got := ClimbStairs(tt.n)
		if got != tt.want {
			t.Errorf("ClimbStairs(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
