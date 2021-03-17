package stack_and_queue

// 维护一个单调减的双端队列，这个队列存可能是最大值的数组下标，这个队列的第一个元素就是最大值下标。
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	for i:=0; i<k; i++ {
		for len(queue)!=0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	res := []int{nums[queue[0]]}
	for i:=k; i<len(nums); i++ {
		for len(queue)!=0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 在这里判断队列第一个元素是否还在窗口里
		for queue[0]+k <= i {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}