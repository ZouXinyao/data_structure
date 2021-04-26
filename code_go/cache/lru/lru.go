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

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		node.value = value
		return
	}
	if this.size == this.capacity {
		this.removeHead()
		this.size--
	}
	newNode := newLinkedNode(key, value)
	this.addToTail(newNode)
	this.cache[key] = newNode
	this.size++
}

func (this *LRUCache) moveToTail(node *LinkedNode) {
	this.removeNode(node)
	this.addToTail(node)
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeHead() {
	node := this.head.next
	this.removeNode(node)
	delete(this.cache, node.key)
}

func (this *LRUCache) addToTail(node *LinkedNode) {
	lastNode := this.tail.prev
	lastNode.next = node
	node.prev = lastNode
	node.next = this.tail
	this.tail.prev = node
}
