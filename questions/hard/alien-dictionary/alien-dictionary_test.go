package aliendictionary_test

import (
	"github.com/jones/interesting-interview-questions/questions/hard/alien-dictionary"
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
		got := aliendictionary.AlienOrder(tt.words)
		if got != tt.want {
			t.Errorf("aliendictionary.AlienOrder(%v) = %q, want %q", tt.words, got, tt.want)
		}
	}
}
