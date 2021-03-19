package backtracking

// 与77类似
func combinationSum(candidates []int, target int) [][]int {
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
		for i := idx; i < n && sum + candidates[i] <= target; i++ {
			sum += candidates[i]
			temp = append(temp, candidates[i])
			helper(i)				// 数字可以重复，所以不需要+1
			sum -= candidates[i]
			temp = temp[:len(temp) - 1]
		}
	}

	helper(0)
	return res
}
