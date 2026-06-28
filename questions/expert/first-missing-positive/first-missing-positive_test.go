package firstmissingpositive_test

import (
	"github.com/jones/interesting-interview-questions/questions/expert/first-missing-positive"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{1}, 2},
		{[]int{}, 1},
	}

	for _, tt := range tests {
		got := firstmissingpositive.FirstMissingPositive(tt.nums)
		if got != tt.want {
			t.Errorf("firstmissingpositive.FirstMissingPositive(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
