package main

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", "c*", true},
	}
	for _, tt := range tests {
		got := IsMatch(tt.s, tt.p)
		if got != tt.want {
			t.Errorf("IsMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
		}
	}
}
