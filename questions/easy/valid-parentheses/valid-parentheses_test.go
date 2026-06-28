package validparentheses_test

import (
	"github.com/jones/interesting-interview-questions/questions/easy/valid-parentheses"
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
		got := validparentheses.IsValid(tt.s)
		if got != tt.want {
			t.Errorf("validparentheses.IsValid(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
