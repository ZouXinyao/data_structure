package main

func subsets(nums []int) [][]int {
	var res [][]int
	var temp []int
	var dfs func(idx int)
	dfs = func(idx int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		// 路径和当前节点都和i有关，idx只不过是起始位置
		for i := idx; i < len(nums); i++ {
			temp = append(temp, nums[i])
			dfs(i + 1)
			temp = temp[:len(temp) - 1]
		}
	}
	dfs(0)
	return res
}
