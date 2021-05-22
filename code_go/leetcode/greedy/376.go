package main

// 题目的用意是有连续上升或连续下降的子集就忽略。只保留上下浮动的。
// 错误理解：有连续res就清零。
// 正确理解：连续时res不变。相等也属于连续的状态。
// 当length为2时，需要两个元素相等，res就是1，其他情况下res为2
func wiggleMaxLength(nums []int) int {
	length := len(nums)
	if length < 2 { return length }

	res := 1
	flag := 0
	for i:=1; i<length; i++ {
		temp := flag
		if nums[i] > nums[i-1] {
			flag = -1
		} else if nums[i] < nums[i-1] {
			flag = 1
		}
		if flag != temp{
			res++
		}
	}
	return res
}
