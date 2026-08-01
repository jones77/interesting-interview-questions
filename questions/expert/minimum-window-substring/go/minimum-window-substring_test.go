package main

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, tt := range tests {
		got := MinWindow(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("MinWindow(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
