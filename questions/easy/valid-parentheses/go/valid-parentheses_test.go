package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"([])", true},
		{"(]", false},
		{"([)]", false},
		{"", false},
		{"(", false},
	}
	for _, tt := range tests {
		got := IsValid(tt.s)
		if got != tt.want {
			t.Errorf("IsValid(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
