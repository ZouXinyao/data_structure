package main

type LFUCache struct {
	size     int
	capacity int
	minFreq  int
	cache    map[int]*LinkedNode
	freqMap  map[int]*LinkedList
}

type LinkedNode struct {
	key, value int
	freq       int
	prev, next *LinkedNode
}

type LinkedList struct {
	size       int
	head, tail *LinkedNode
}

func newLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

func newLinkedList() *LinkedList {
	l := &LinkedList{
		size: 0,
		head: newLinkedNode(0, 0),
		tail: newLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func Constructor(capacity int) LFUCache {
	c := LFUCache{
		capacity: capacity,
		cache:    map[int]*LinkedNode{},
		freqMap:  map[int]*LinkedList{},
	}
	return c
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.updateFreq(node)
		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key, value int) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.cache[key]; ok {
		this.updateFreq(node)
		node.value = value
		return
	}
	if this.size == this.capacity {
		this.removeMinFreqNode()
		this.size--
	}
	newNode := newLinkedNode(key, value)
	this.updateFreq(newNode)
	this.cache[key] = newNode
	this.size++
}

func (this *LFUCache) updateFreq(node *LinkedNode) {
	f := node.freq
	if f > 0 {
		this.freqMap[f].removeNode(node)
		this.freqMap[f].size--
		if this.freqMap[f].size == 0 {
			if this.minFreq == f {
				this.minFreq++
			}
			delete(this.freqMap, f)
		}
	} else {
		this.minFreq = 1
	}
	node.freq++
	if _, ok := this.freqMap[node.freq]; !ok {
		this.freqMap[node.freq] = newLinkedList()
	}
	this.freqMap[node.freq].addToTail(node)
	this.freqMap[node.freq].size++
}

func (this *LFUCache) removeMinFreqNode() {
	node := this.freqMap[this.minFreq].removeHead()
	delete(this.cache, node.key)
	this.freqMap[this.minFreq].size--
	if this.freqMap[this.minFreq].size == 0 {
		delete(this.freqMap, this.minFreq)
	}
}

func (this *LinkedList) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LinkedList) addToTail(node *LinkedNode) {
	lastNode := this.tail.prev
	lastNode.next = node
	node.prev = lastNode
	node.next = this.tail
	this.tail.prev = node
}

func (this *LinkedList) removeHead() *LinkedNode {
	node := this.head.next
	this.removeNode(node)
	return node
}

func main() {
	lfu := Constructor(0)
	lfu.Put(0, 0)
	lfu.Get(0)
}
