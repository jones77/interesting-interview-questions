package reversestring_test

import (
	"github.com/jones/interesting-interview-questions/questions/easy/reverse-string"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
	}
	for _, tt := range tests {
		reversestring.ReverseString(tt.s)
		if !reflect.DeepEqual(tt.s, tt.want) {
			t.Errorf("reverseString = %v, want %v", tt.s, tt.want)
		}
	}
}
