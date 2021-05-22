package stack_and_queue

// 维护一个单调减的 双端队列 ，这个队列存的是 可能是最大值 的数组下标，这个队列的第一个元素就是最大值下标。
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	for i := 0; i < k; i++ {
		// 新元素比尾元素大，就弹出
		for len(queue) != 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	res := []int{nums[queue[0]]}
	for i := k; i < len(nums); i++ {
		for len(queue) != 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 在这里判断队列第一个元素是否还在窗口里
		for queue[0]+k <= i {
			queue = queue[1:]
		}
		// 队列中的首元素一定是最大的，而且在窗里，所以每移动一次窗，就要往res中添加一个结果。
		res = append(res, nums[queue[0]])
	}
	return res
}
