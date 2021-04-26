package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	lastRight := intervals[0][1]
	// 按尾排序后，不需要关心头在哪里；
	// 可以将排序后的每不相邻的区间看做各个集合；
	// 当a与b重合，b与c重合，a与c不重合，去掉b即可，a和c满足不重合的条件就是a的尾<=c的头；
	// 所以，a的尾<=c的头时，不需要去掉区间，直接将c的尾部作为tempRight；
	// a的尾>b的头时，需要去掉b区间，这里就直接记录需要去掉的数量，而且不改变tempRight。
	// 一直遍历到最后一个元素就可以了
	for i:=1; i<len(intervals); i++ {
		if intervals[i][0] >= lastRight {
			lastRight = intervals[i][1]
		} else {
			res++
		}
	}
	return res
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(eraseOverlapIntervals(points))
}
