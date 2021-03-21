package backtracking

import "sort"

// 去重就需要排序+条件判断，这个条件一般都是相邻的两个数相等的话，后面那个数就跳过了
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var temp []int
	var dfs func(idx int)
	dfs = func(idx int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := idx; i < len(nums); i++ {
			if i != idx && nums[i] == nums[i - 1] { continue }
			temp = append(temp, nums[i])
			dfs(i + 1)
			temp = temp[:len(temp) - 1]
		}
	}
	dfs(0)
	return res
}
