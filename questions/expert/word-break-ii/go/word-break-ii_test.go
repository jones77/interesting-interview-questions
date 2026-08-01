package main

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     []string
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{"cats and dog", "cat sand dog"}},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}, []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, nil},
	}

	for _, tt := range tests {
		got := WordBreak(tt.s, tt.wordDict)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("WordBreak(%v, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
		}
	}
}
