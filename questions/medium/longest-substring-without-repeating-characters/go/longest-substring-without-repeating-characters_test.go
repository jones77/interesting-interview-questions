package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
	}
	for _, tt := range tests {
		got := LengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
