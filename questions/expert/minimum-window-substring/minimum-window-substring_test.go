package minimumwindowsubstring_test

import (
	"github.com/jones/interesting-interview-questions/questions/expert/minimum-window-substring"
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
		got := minimumwindowsubstring.MinWindow(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("minimumwindowsubstring.MinWindow(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
