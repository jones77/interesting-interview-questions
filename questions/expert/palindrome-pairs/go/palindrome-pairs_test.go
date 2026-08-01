package main

import (
	"reflect"
	"testing"
)

func TestPalindromePairs(t *testing.T) {
	tests := []struct {
		words []string
		want  [][]int
	}{
		{[]string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
		{[]string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}}},
		{[]string{"a", ""}, [][]int{{0, 1}, {1, 0}}},
	}

	for _, tt := range tests {
		got := PalindromePairs(tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PalindromePairs(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}
}
