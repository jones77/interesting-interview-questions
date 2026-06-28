package main

import (
	"reflect"
	"testing"
)

// Implement: func productExceptSelf(nums []int) []int

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		got := productExceptSelf(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
