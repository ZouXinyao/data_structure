package dp

import "math"

func maxProfit(prices []int) int {
	low := math.MaxInt64
	res := 0
	for _, p := range prices {
		low = min(low, p)
		res = max(res, p - low)
	}
	return res
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func max(a, b int) int {
	if a < b { return b }
	return a
}