package fizzbuzz_test

import (
	"github.com/jones/interesting-interview-questions/questions/easy/fizzbuzz"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		got := fizzbuzz.FizzBuzz(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("fizzbuzz.FizzBuzz(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
