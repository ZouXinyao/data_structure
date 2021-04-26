package stack_and_queue

type MyCircularDeque struct {
	queue []int
	first int
	rear  int
	size  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k, k),
		first: 0,
		rear:  1,
		size:  0,
	}
}

// 添加时，先添加，再移位
/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() { return false }
	this.queue[this.first] = value
	this.size++
	this.first = (len(this.queue) + this.first-1) % (len(this.queue))
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() { return false }
	this.queue[this.rear] = value
	this.size++
	this.rear = (this.rear+1) % (len(this.queue))
	return true
}

// 删除时，先移动索引，再删除
/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() { return false }
	this.first = (this.first+1) % (len(this.queue))
	this.queue[this.first] = 0
	this.size--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() { return false }
	this.rear = (len(this.queue)+this.rear-1) % (len(this.queue))
	this.queue[this.rear] = 0
	this.size--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() { return -1 }
	return this.queue[(this.first+1) % (len(this.queue))]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() { return -1 }
	return this.queue[(len(this.queue)+this.rear-1) % (len(this.queue))]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.size == len(this.queue)
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
