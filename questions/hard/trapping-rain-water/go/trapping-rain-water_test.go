package main

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{}, 0},
		{[]int{1}, 0},
	}
	for _, tt := range tests {
		got := Trap(tt.height)
		if got != tt.want {
			t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
		}
	}
}
