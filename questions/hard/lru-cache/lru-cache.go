// LRU Cache
//
// Design a data structure that follows the constraints of a Least Recently
// Used (LRU) cache.
//
// Implement the LRUCache class:
//   - LRUCache(capacity) initialise with positive size capacity
//   - int get(key) return the value if the key exists, otherwise -1
//   - void put(key, value) update or insert the key-value pair; evict the
//     least recently used key if the number of keys exceeds capacity
//
// Both get and put must run in O(1) average time.
//
// Example:
//   Input
//   ["LRUCache","put","put","get","put","get","put","get","get","get"]
//   [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
//   Output
//   [null,null,null,1,null,-1,null,-1,3,4]

package main

type LRUCache struct{}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (c *LRUCache) Get(key int) int {
	return -1
}

func (c *LRUCache) Put(key int, value int) {
}

func main() {}
