package main

import (
	"sort"
)
// 从高到低排列，然后按条件拼接数组
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	var res [][]int
	for _, p := range people {
		idx := p[1]
		res = append(res[:idx], append([][]int{p}, res[idx:]...)...)
	}

	return res
}