package stack_and_queue

// 单调减的栈。
func trap(height []int) int {
	res := 0
	stack := []int{}
	for i:=0; i < len(height); i++ {
		for len(stack) != 0 && height[stack[len(stack) - 1]] < height[i] {
			cur := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 { break }
			// 计算左右边界
			l, r := stack[len(stack) - 1], i
			// 计算雨滴
			h := min(height[r], height[l]) - height[cur]
			res += (r - l - 1) * h
		}
		stack = append(stack, i)
	}
	return res
}

func min(a, b int) int {
	if a < b { return a }
	return b
}
