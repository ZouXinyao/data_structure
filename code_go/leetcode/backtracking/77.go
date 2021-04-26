package main

func combine(n int, k int) [][]int {
	var res  [][]int
	var temp []int
	var helper func(int)
	helper = func(idx int) {
		if len(temp) == k {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		// n - (k - len(temp)) + 1用来剪枝；普通的回溯法是：i <= n
		for i := idx; i <= n - (k - len(temp)) + 1; i++ {
			temp = append(temp, i)
			helper(i + 1)
			temp = temp[:len(temp) - 1]
		}
	}

	helper(1)
	return res
}