package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func newLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		cache:    map[int]*LinkedNode{},
		head:     newLinkedNode(0, 0),
		tail:     newLinkedNode(0, 0),
	}

	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToTail(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.cache[key]; ok {
		lru.moveToTail(node)
		node.value = value
		return
	}
	if lru.size == lru.capacity {
		lru.removeHead()
		lru.size--
	}
	newNode := newLinkedNode(key, value)
	lru.addToTail(newNode)
	lru.cache[key] = newNode
	lru.size++
}

func (lru *LRUCache) moveToTail(node *LinkedNode) {
	lru.removeNode(node)
	lru.addToTail(node)
}

func (lru *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) removeHead() {
	node := lru.head.next
	lru.removeNode(node)
	delete(lru.cache, node.key)
}

func (lru *LRUCache) addToTail(node *LinkedNode) {
	lastNode := lru.tail.prev
	lastNode.next = node
	node.prev = lastNode
	node.next = lru.tail
	lru.tail.prev = node
}
