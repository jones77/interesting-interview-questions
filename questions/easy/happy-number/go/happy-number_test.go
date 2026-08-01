package main

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
		{1, true},
	}
	for _, tt := range tests {
		got := IsHappy(tt.n)
		if got != tt.want {
			t.Errorf("IsHappy(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
