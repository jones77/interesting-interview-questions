package testutil

import (
	"reflect"
	"sort"
	"testing"
)

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertEqualSliceInt(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("len(got)=%d, len(want)=%d", len(got), len(want))
		return
	}
	g := append([]int(nil), got...)
	w := append([]int(nil), want...)
	sort.Ints(g)
	sort.Ints(w)
	if !reflect.DeepEqual(g, w) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertEqualSliceString(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("len(got)=%d, len(want)=%d", len(got), len(want))
		return
	}
	g := append([]string(nil), got...)
	w := append([]string(nil), want...)
	sort.Strings(g)
	sort.Strings(w)
	if !reflect.DeepEqual(g, w) {
		t.Errorf("got %v, want %v", got, want)
	}
}
