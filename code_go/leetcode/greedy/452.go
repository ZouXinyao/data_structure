package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func (i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	maxRight := points[0][1]

	// 按尾排序后，箭射在尾部；
	// 从左到右根据尾部找有间隔的气球；
	// 当头在前面最短尾的后面，res++，当前气球更新为第1个，也就是最短尾。
	for i:=1; i<len(points); i++ {
		if points[i][0] > maxRight {
			maxRight = points[i][1]
			res++
		}
	}
	return res
}

//func main() {
//	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
//	fmt.Println(findMinArrowShots(points))
//}
