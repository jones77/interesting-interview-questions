package main

import (
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
		reverseString(tt.s)
		if !reflect.DeepEqual(tt.s, tt.want) {
			t.Errorf("reverseString = %v, want %v", tt.s, tt.want)
		}
	}
}
