package palindromepairs_test

import (
	"github.com/jones/interesting-interview-questions/questions/expert/palindrome-pairs"
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
		got := palindromepairs.PalindromePairs(tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("palindromepairs.PalindromePairs(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}
}
