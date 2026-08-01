package main

import (
	"reflect"
	"sort"
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
		got := TwoSum(tt.nums, tt.target)
		g := append([]int(nil), got...)
		sort.Ints(g)
		w := append([]int(nil), tt.want...)
		sort.Ints(w)
		if !reflect.DeepEqual(g, w) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
