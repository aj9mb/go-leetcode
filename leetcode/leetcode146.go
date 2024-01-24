package leetcode

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type LRUCache struct {
	cap   int
	cache map[int]*node
	left  *node
	right *node
}

func Construct(capacity int) LRUCache {
	left := &node{key: 0, value: 0}
	right := &node{key: 0, value: 0}
	left.next, right.prev = right, left
	return LRUCache{cap: capacity, cache: make(map[int]*node), left: left, right: right}
}

func remove(n *node) {
	prev, nxt := n.prev, n.next
	prev.next, nxt.prev = nxt, prev
}

func insert(lru *LRUCache, n *node) {
	prev, nxt := lru.right.prev, lru.right
	prev.next, nxt.prev = n, n
	n.prev, n.next = prev, nxt
}

func (cache *LRUCache) Get(key int) int {
	if n, ok := cache.cache[key]; ok {
		remove(n)
		insert(cache, n)
		return n.value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if n, ok := cache.cache[key]; ok {
		remove(n)
	}
	n := &node{key: key, value: value}
	cache.cache[key] = n
	insert(cache, n)
	if len(cache.cache) > cache.cap {
		delete(cache.cache, cache.left.next.key)
		remove(cache.left.next)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
