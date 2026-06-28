// Copy this file and a solution stub into the question directory:
//
//   cp skeleton_test.go <short-name>_test.go
//   cp skeleton_solution.go <short-name>.go
//
// Then implement the function in <short-name>.go
// and run tests with: go test -v ./questions/<difficulty>/<short-name>/

package main

import (
	"reflect"
	"testing"
)

func TestQuestion(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want any
	}{
		// Add expected inputs and outputs here.
		// {"example 1", []any{...}, ...},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := function(tt.args...)
			// if !reflect.DeepEqual(got, tt.want) {
			//     t.Errorf("got %v, want %v", got, tt.want)
			// }
		})
	}
}
