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

func (this *LRUCache) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		remove(n)
		insert(this, n)
		return n.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.cache[key]; ok {
		remove(n)
	}
	n := &node{key: key, value: value}
	this.cache[key] = n
	insert(this, n)
	if len(this.cache) > this.cap {
		delete(this.cache, this.left.next.key)
		remove(this.left.next)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
