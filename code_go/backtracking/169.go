package main

func majorityElement(nums []int) int {
	var count func(num, l, r int) int			// 求num在l-r范围内的数量
	var majorityElementRec func(l, r int) int	// 求最多的num
	count = func(num, l, r int) int {
		var c int
		for i := l; i <= r; i++ {
			if nums[i] == num { c++ }
		}
		return c
	}
	// 最多的数一定是在左半最多或者右半最多。
	majorityElementRec = func(l, r int) int {
		if l == r {
			return nums[l]
		}

		m := (r - l) / 2 + l
		// 求左半最多和右半最多，然后返回最多的那个数。
		left := majorityElementRec(l, m)
		right := majorityElementRec(m + 1, r)

		if left == right { return left }

		leftCount := count(left, l, m)
		rightCount := count(right, m + 1, r)
		if leftCount > rightCount { return left }
		return right
	}
	return majorityElementRec(0, len(nums) - 1)
}
