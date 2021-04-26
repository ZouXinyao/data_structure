package stack_and_queue

import "math"

type MinStack struct {
	size     int
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		size: 0,
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	this.minStack = append(this.minStack, min(x, this.minStack[this.size]))
	this.size++
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:this.size - 1]
	this.minStack = this.minStack[:this.size]
	this.size--
}


func (this *MinStack) Top() int {
	return  this.stack[this.size - 1]
}


func (this *MinStack) GetMin() int {
	return  this.minStack[this.size]
}

func min(a, b int) int {
	if a > b { return b }
	return a
}
