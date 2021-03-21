package main

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(temp []byte, l, r, target int)
	dfs = func(temp []byte, l, r, target int) {
		if l == target && r == target {
			res = append(res, string(temp))
			return
		}

		if l < target {
			dfs(append(temp, '('), l + 1, r, target)
		}
		if r < l {
			dfs(append(temp, ')'), l, r + 1, target)
		}
	}
	dfs([]byte{}, 0, 0, n)
	return res
}
