package backtracking

func partition(s string) [][]string {
	var res [][]string
	var temp []string
	var palindrome func(str string, l, r int) bool
	var dfs func(idx int)
	palindrome = func(str string, l, r int) bool {
		for l < r {
			if str[l] == str[r] {
				l++
				r--
			} else {
				return false
			}
		}
		return true
	}
	dfs = func(idx int) {
		if idx == len(s) {                  // 当start指针越界了，一直切出回文才走到这，将当前的部分解temp加入解集res
			t := make([]string, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := idx; i < len(s); i++ {             // 枚举出当前的所有选项，从索引idx到末尾索引
			if palindrome(s, idx, i) {              // 当前选择i，如果 idx 到 i 是回文串，就添加到temp
				temp = append(temp, s[idx : i+1])
				dfs(i + 1)                          // 基于这个选择，继续往下递归，继续
				temp = temp[: len(temp) - 1]        // 上面递归结束了，撤销当前选择i，去下一轮迭代
			}
		}
	}

	dfs(0)
	return res
}
