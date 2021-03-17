package main

import "fmt"

// 暴力解法
//func largestRectangleArea(heights []int) int {
//	l, r := 0, 0
//	res := 0
//	for i:=0; i<len(heights); i++ {
//		for j:=i; j<len(heights); j++ {
//			if heights[j] < heights[i]{
//				r = j - 1
//				break
//			}
//			r = j
//		}
//		for k:=i; k>=0; k-- {
//			if heights[k] < heights[i] {
//				l = k + 1
//				break
//			}
//			l = k
//		}
//		res = max(res, heights[i] * (r - l + 1))
//	}
//	return res
//}
//
//func max(a, b int) int {
//	if a<b { return b }
//	return a
//}

// 单调栈，维护一个单调增的栈，用来存放右边界之前的所有值
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := [][]int{{-1, 0}}
	res := 0
	for i:=0; i<len(heights); i++ {
		l, r := i - 1, i - 1
		for heights[i] < stack[len(stack) - 1][1] {
			// -2是为了取上一个栈中前面那个元素的位置，因为上一个位置才是边界。
			l = stack[len(stack) - 2][0]
			res = max(res, stack[len(stack) - 1][1] * (r - l))
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, []int{i, heights[i]})
	}
	return res
}

func max(a, b int) int {
	if a<b { return b }
	return a
}

func main() {
	height := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(height))
}