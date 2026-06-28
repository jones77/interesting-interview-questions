package maximumsubarray_test

import (
	"github.com/jones/interesting-interview-questions/questions/easy/maximum-subarray"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1}, -1},
	}
	for _, tt := range tests {
		got := maximumsubarray.MaxSubArray(tt.nums)
		if got != tt.want {
			t.Errorf("maximumsubarray.MaxSubArray(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
