package findduplicatenumber_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/find-duplicate-number"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 1}, 1},
	}
	for _, tt := range tests {
		got := findduplicatenumber.FindDuplicate(tt.nums)
		if got != tt.want {
			t.Errorf("findduplicatenumber.FindDuplicate(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
