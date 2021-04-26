package list

// 双指针
func maxArea(height []int) int {
	l, r, res := 0, len(height) - 1, 0
	for l < r {
		res = max(res, min(height[l], height[r]) * (r - l))
		// 判断条件是哪边矮，移哪边
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a>b {return a}
	return b
}

func min(a, b int) int {
	if a<b {return a}
	return b
}
