package main

import "fmt"

type FIFOCache struct {
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

func Constructor(capacity int) FIFOCache {
	c := FIFOCache{
		capacity: capacity,
		cache:    map[int]*LinkedNode{},
		head:     newLinkedNode(0, 0),
		tail:     newLinkedNode(0, 0),
	}

	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (this *FIFOCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		return node.value
	}
	return -1
}

func (this *FIFOCache) Put(key, value int) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.cache[key]; ok {
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

func (this *FIFOCache) removeHead() {
	node := this.head.next
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(this.cache, node.key)
}

func (this *FIFOCache) addToTail(node *LinkedNode) {
	lastNode := this.tail.prev
	lastNode.next = node
	node.prev = lastNode
	node.next = this.tail
	this.tail.prev = node
}

func main() {
	fifo := Constructor(2)
	fifo.Put(0, 0)
	fifo.Put(1, 1)
	fmt.Println(fifo.Get(0))
	fifo.Put(2, 2)
	fmt.Println(fifo.Get(0))
}
