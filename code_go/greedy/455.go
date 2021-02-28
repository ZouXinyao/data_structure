package main

import "sort"

func findContentChildren(children []int, cookies []int) int {
	sort.Ints(children)
	sort.Ints(cookies)

	res := 0
	for i, j := 0, 0; i < len(children) && j < len(cookies); j++ {
		if children[i] <= cookies[j] {
			i++
			res++
		}
	}
	return res
}