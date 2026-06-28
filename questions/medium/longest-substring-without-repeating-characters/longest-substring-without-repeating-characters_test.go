package longestsubstringwithoutrepeatingcharacters_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/longest-substring-without-repeating-characters"
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
		got := longestsubstringwithoutrepeatingcharacters.LengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("longestsubstringwithoutrepeatingcharacters.LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
