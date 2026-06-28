package evaluatereversepolishnotation_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/evaluate-reverse-polish-notation"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		want   int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, tt := range tests {
		got := evaluatereversepolishnotation.EvalRPN(tt.tokens)
		if got != tt.want {
			t.Errorf("evaluatereversepolishnotation.EvalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
		}
	}
}
