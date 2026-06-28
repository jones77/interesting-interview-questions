package twosum_test

import (
	"github.com/jones/interesting-interview-questions/internal/testutil"
	"github.com/jones/interesting-interview-questions/questions/easy/two-sum"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		got := twosum.TwoSum(tt.nums, tt.target)
		testutil.AssertEqualSliceInt(t, got, tt.want)
	}
}
