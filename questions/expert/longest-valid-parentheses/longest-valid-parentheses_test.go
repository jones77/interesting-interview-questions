package longestvalidparentheses_test

import (
	"github.com/jones/interesting-interview-questions/questions/expert/longest-valid-parentheses"
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
		got := longestvalidparentheses.LongestValidParentheses(tt.s)
		if got != tt.want {
			t.Errorf("longestvalidparentheses.LongestValidParentheses(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
