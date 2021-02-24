package main

import (
	"sort"
)

func partitionLabels(S string) []int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	byt := []byte(S)

	// 用map存下每个字符的出现第一次的索引和最后依次的索引
	tempMap := map[byte][]int{}
	for i, b := range byt {
		if _, ok := tempMap[b]; ok {
			tempMap[b][1] = i
		} else {
			tempMap[b] = []int{i, i}
		}
	}

	// 转存到切片中，将该问题转换为求每个相隔集合长度
	tempSlice := [][]int{}
	for _, m := range tempMap {
		tempSlice = append(tempSlice, m)
	}

	// 切片按左值排序
	sort.Slice(tempSlice, func(i, j int) bool {
		return tempSlice[i][0] < tempSlice[j][0]
	})

	// 根据求每个相交集合的最左值和最右值，求出每个集合的长度。
	res := []int{}
	left, right := tempSlice[0][0], tempSlice[0][1]
	for i := 1; i < len(tempSlice); i++ {
		if tempSlice[i][0] >= right {
			res = append(res, right-left+1)
			left, right = tempSlice[i][0], tempSlice[i][1]
		} else {
			right = max(tempSlice[i][1], right)
		}
	}
	res = append(res, right-left+1)
	return res
}

//func main() {
//	S := "ababcbacadefegdehijhklij"
//	fmt.Println(partitionLabels(S))
//}
