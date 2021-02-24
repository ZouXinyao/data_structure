package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	max := func (a, b int) int {
		if a > b { return a }
		return b
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res   := [][]int{}
	left  := intervals[0][0]
	right := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			// 当没有重叠时，我们就收获了1个子区间，并且要更新left, right作为新子区间的初始值
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			// 有重叠时，只需要将right更新到最右边就可以了。
			right = max(right, intervals[i][1])
		}
	}
	res = append(res, []int{left, right})
	return res
}



//func main() {
//	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
//	fmt.Println(merge(points))
//}
