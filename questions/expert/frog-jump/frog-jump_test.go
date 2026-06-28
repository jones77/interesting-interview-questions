package frogjump_test

import (
	"github.com/jones/interesting-interview-questions/questions/expert/frog-jump"
	"testing"
)

func TestCanCross(t *testing.T) {
	tests := []struct {
		stones []int
		want   bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 1}, true},
	}

	for _, tt := range tests {
		got := frogjump.CanCross(tt.stones)
		if got != tt.want {
			t.Errorf("frogjump.CanCross(%v) = %v, want %v", tt.stones, got, tt.want)
		}
	}
}
