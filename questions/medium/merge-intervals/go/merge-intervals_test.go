package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}}, [][]int{{1, 4}}},
	}
	for _, tt := range tests {
		got := Merge(tt.intervals)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Merge(%v) = %v, want %v", tt.intervals, got, tt.want)
		}
	}
}
