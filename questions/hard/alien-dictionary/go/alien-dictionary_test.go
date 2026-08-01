package main

import (
	"testing"
)

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{[]string{"z", "x", "z"}, ""},
		{[]string{"z", "x"}, "zx"},
	}
	for _, tt := range tests {
		got := AlienOrder(tt.words)
		if got != tt.want {
			t.Errorf("AlienOrder(%v) = %q, want %q", tt.words, got, tt.want)
		}
	}
}
