package leetcode

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Standard sequence from LeetCode", func(t *testing.T) {
		// Initialize cache with capacity 2
		lru := Constructor(2)

		lru.Put(1, 1) // cache is {1=1}
		lru.Put(2, 2) // cache is {1=1, 2=2}

		if got := lru.Get(1); got != 1 {
			t.Errorf("Get(1) = %d; want 1", got)
		}

		lru.Put(3, 3) // LRU was key 2, evicts key 2, cache is {1=1, 3=3}

		if got := lru.Get(2); got != -1 {
			t.Errorf("Get(2) = %d; want -1 (evicted)", got)
		}

		lru.Put(4, 4) // LRU was key 1, evicts key 1, cache is {4=4, 3=3}

		if got := lru.Get(1); got != -1 {
			t.Errorf("Get(1) = %d; want -1 (evicted)", got)
		}

		if got := lru.Get(3); got != 3 {
			t.Errorf("Get(3) = %d; want 3", got)
		}

		if got := lru.Get(4); got != 4 {
			t.Errorf("Get(4) = %d; want 4", got)
		}
	})

	t.Run("Update existing key", func(t *testing.T) {
		lru := Constructor(1)
		lru.Put(2, 1)
		if got := lru.Get(2); got != 1 {
			t.Errorf("Get(2) = %d; want 1", got)
		}

		// Update value for key 2
		lru.Put(2, 2)
		if got := lru.Get(2); got != 2 {
			t.Errorf("Get(2) after update = %d; want 2", got)
		}
	})
}
