package candy_test

import (
	"github.com/jones/interesting-interview-questions/questions/expert/candy"
	"testing"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1}, 1},
		{[]int{1, 3, 2, 2, 1}, 7},
	}

	for _, tt := range tests {
		got := candy.Candy(tt.ratings)
		if got != tt.want {
			t.Errorf("candy.Candy(%v) = %v, want %v", tt.ratings, got, tt.want)
		}
	}
}
