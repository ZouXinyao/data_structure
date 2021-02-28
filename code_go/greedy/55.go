package main

// 以保存的最大跳跃范围为贪心条件。
func canJump(nums []int) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxJump := 0
	for i, k := range nums {
		maxJump = max(i+k, maxJump)
		if i >= maxJump {
			return false
		}
		if maxJump >= len(nums)-1 {
			return true
		}
	}
	return true
}
