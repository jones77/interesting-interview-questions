package climbingstairs_test

import (
	"github.com/jones/interesting-interview-questions/questions/easy/climbing-stairs"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, tt := range tests {
		got := climbingstairs.ClimbStairs(tt.n)
		if got != tt.want {
			t.Errorf("climbingstairs.ClimbStairs(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
