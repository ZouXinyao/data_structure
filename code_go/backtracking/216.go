package main

func combinationSum3(k int, n int) [][]int {
	sum := 0
	var res  [][]int
	var temp []int
	var helper func(int)
	helper = func(idx int) {
		if len(temp) > k || sum > n { return }
		if len(temp) == k && sum == n {
			t := make([]int, k)
			copy(t, temp)
			res = append(res, t)
			return
		}

		for i := idx; i <= 9; i++ {
			sum += i
			temp = append(temp, i)
			helper(i + 1)
			sum -= i
			temp = temp[:len(temp) - 1]
		}

	}
	helper(1)
	return res
}
