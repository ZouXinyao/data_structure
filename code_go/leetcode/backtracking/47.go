package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var temp []int
	var dfs func()
	used := make([]bool, len(nums))
	dfs = func() {
		if len(temp) == len(nums) {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
		}
		for i := 0; i < len(nums); i++ {
			// 去重条件：!used[i - 1]同一层的前一位读过了，所以这次不能读。
			// if i != 0 && nums[i] == nums[i - 1] && used[i - 1] { continue } 也可以
			if i != 0 && nums[i] == nums[i - 1] && !used[i - 1] { continue }
			if used[i] { continue }
			used[i] = true
			temp = append(temp, nums[i])
			dfs()
			temp = temp[:len(temp) - 1]
			used[i] = false
		}
	}
	dfs()
	return res
}
