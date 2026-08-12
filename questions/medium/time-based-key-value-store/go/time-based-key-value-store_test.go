package main

import (
	"testing"
)

func TestTimeMap(t *testing.T) {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	if got := tm.Get("foo", 1); got != "bar" {
		t.Errorf("Get(foo, 1) = %q, want bar", got)
	}
	if got := tm.Get("foo", 3); got != "bar" {
		t.Errorf("Get(foo, 3) = %q, want bar", got)
	}
	tm.Set("foo", "bar2", 4)
	if got := tm.Get("foo", 4); got != "bar2" {
		t.Errorf("Get(foo, 4) = %q, want bar2", got)
	}
	if got := tm.Get("foo", 5); got != "bar2" {
		t.Errorf("Get(foo, 5) = %q, want bar2", got)
	}
}
