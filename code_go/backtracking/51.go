package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}

	// 判断当前位置是否符合要求，
	var isValid func(r, c int, bd [][]string) bool
	isValid = func(r, c int, bd [][]string) bool {
		for i := 0; i < r; i++ {
			for j := 0; j < n; j++ {
				// 列上、两个斜线上是否有Q
				if bd[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
					return false
				}
			}
		}
		return true
	}

	res := [][]string{}
	var helper func(r int, bd [][]string)
	helper = func(r int, bd [][]string) {
		if r == n {
			temp := make([]string, len(bd))
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(bd[i], "")
			}
			res = append(res, temp)
			return
		}
		for c := 0; c < n; c++ {
			// 只有满足要求，才能继续往下走
			if isValid(r, c, bd) {
				bd[r][c] = "Q"
				helper(r+1, bd)
				bd[r][c] = "."
			}
		}
	}

	helper(0, bd)
	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
