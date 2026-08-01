package main

import (
	"reflect"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1, -1}, []int{0, 0}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		got := CountSmaller(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CountSmaller(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
