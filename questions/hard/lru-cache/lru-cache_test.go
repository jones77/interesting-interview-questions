package lrucache_test

import (
	"github.com/jones/interesting-interview-questions/questions/hard/lru-cache"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := lrucache.Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if got := cache.Get(1); got != 1 {
		t.Errorf("lrucache.Get(1) = %d, want 1", got)
	}
	cache.Put(3, 3)
	if got := cache.Get(2); got != -1 {
		t.Errorf("lrucache.Get(2) = %d, want -1 (evicted)", got)
	}
	cache.Put(4, 4)
	if got := cache.Get(1); got != -1 {
		t.Errorf("lrucache.Get(1) = %d, want -1 (evicted)", got)
	}
	if got := cache.Get(3); got != 3 {
		t.Errorf("lrucache.Get(3) = %d, want 3", got)
	}
	if got := cache.Get(4); got != 4 {
		t.Errorf("lrucache.Get(4) = %d, want 4", got)
	}
}
