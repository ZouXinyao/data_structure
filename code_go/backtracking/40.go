package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	sum := 0
	var res  [][]int
	var temp []int
	var helper func(int)
	helper = func(idx int) {
		if sum > target { return }
		if sum == target {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := idx; i < n; i++ {
			// 排序后避免重复，需要当前元素和他前一个元素不等就行了
			if i - 1 >= idx && candidates[i] == candidates[i-1] { continue }

			sum += candidates[i]
			temp = append(temp, candidates[i])
			helper(i+1)				// 数字可以重复，所以不需要+1
			sum -= candidates[i]
			temp = temp[:len(temp) - 1]
		}
	}
	sort.Ints(candidates)
	helper(0)
	return res
}
