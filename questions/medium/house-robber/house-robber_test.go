package houserobber_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/house-robber"
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1}, 2},
		{[]int{1}, 1},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		got := houserobber.Rob(tt.nums)
		if got != tt.want {
			t.Errorf("houserobber.Rob(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
