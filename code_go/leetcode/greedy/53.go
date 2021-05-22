package main

import "math"

// 贪心解法，核心思路是从头开始序列和sum,sum<0就不考虑，sum>0时保留;实时保存结果res，当sum>res就要这个结果
func maxSubArray(nums []int) int {
	result := math.MinInt32
	tempSum := 0
	for _, k := range nums {
		tempSum += k
		if tempSum > result { result = tempSum }
		if tempSum < 0 { tempSum = 0 }
	}
	return result
}
