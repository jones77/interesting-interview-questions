package main

import (
	"testing"
)

func TestCanCross(t *testing.T) {
	tests := []struct {
		stones []int
		want   bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 1}, true},
	}

	for _, tt := range tests {
		got := CanCross(tt.stones)
		if got != tt.want {
			t.Errorf("CanCross(%v) = %v, want %v", tt.stones, got, tt.want)
		}
	}
}
