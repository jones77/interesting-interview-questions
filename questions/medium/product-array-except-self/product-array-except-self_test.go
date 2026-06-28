package productarrayexceptself_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/product-array-except-self"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		got := productarrayexceptself.ProductExceptSelf(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("productarrayexceptself.ProductExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
