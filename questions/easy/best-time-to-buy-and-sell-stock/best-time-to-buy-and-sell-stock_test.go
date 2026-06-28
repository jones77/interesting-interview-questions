package besttimetobuyandsellstock_test

import (
	"github.com/jones/interesting-interview-questions/questions/easy/best-time-to-buy-and-sell-stock"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
	}
	for _, tt := range tests {
		got := besttimetobuyandsellstock.MaxProfit(tt.prices)
		if got != tt.want {
			t.Errorf("besttimetobuyandsellstock.MaxProfit(%v) = %v, want %v", tt.prices, got, tt.want)
		}
	}
}
