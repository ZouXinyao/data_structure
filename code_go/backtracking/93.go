package backtracking

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string
	var temp []string
	var dfs func(idx int)
	dfs = func(idx int) {
		// 终止条件是temp有4段，并且idx也到了s的结尾
		if len(temp) == 4 && idx == len(s) {
			res = append(res, strings.Join(temp, "."))
			return
		}
		if len(temp) == 4 && idx < len(s) {
			return
		}
		for l := 1; l <= 3; l++ {
			// 有几种情况是需要直接返回：1、idx+l超范围；2、当前段开头为0
			if idx + l - 1 >= len(s) { return }
			if l != 1 && s[idx] == '0' { return }
			str := s[idx : l + idx]
			n, _ := strconv.Atoi(str);
			if n > 255 { return }
			temp = append(temp, str)
			dfs(idx + l)
			temp = temp[:len(temp) - 1]
		}
	}

	dfs(0)
	return res
}
