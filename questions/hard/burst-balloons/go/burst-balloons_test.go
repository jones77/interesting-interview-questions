package main

import (
	"testing"
)

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		got := MaxCoins(tt.nums)
		if got != tt.want {
			t.Errorf("MaxCoins(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
