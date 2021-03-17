package list

// 3次换序
func rotate(nums []int, k int)  {
	l, r := 0, len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		r--
		l++
	}
	l, r = 0, (k % len(nums)) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		r--
		l++
	}
	l, r = k % len(nums), len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		r--
		l++
	}
}
