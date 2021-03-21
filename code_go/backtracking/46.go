package backtracking

func permute(nums []int) [][]int {
	var res [][]int
	var temp []int
	var dfs func()
	// 用1个数组存一下当前这个位置的数字用没用过，用过就跳过了。
	used := make([]bool, len(nums))
	dfs = func() {
		if len(temp) == len(nums) {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
		}
		for i := 0; i < len(nums); i++ {
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
