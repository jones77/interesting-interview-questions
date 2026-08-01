package main

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"()(()", 2},
		{"()(())", 6},
	}

	for _, tt := range tests {
		got := LongestValidParentheses(tt.s)
		if got != tt.want {
			t.Errorf("LongestValidParentheses(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
