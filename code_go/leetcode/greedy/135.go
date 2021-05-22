package main

func candy(rating []int) int {
	num := len(rating)
	if num < 2 {
		return num
	}
	// 两次贪心，一次向后遍历，从1开始，也就是中和左相比；另一次是向前，从length-2开始，中和右相比
	candyList := make([]int, num)
	for i := 1; i < num; i++ {
		if rating[i] > rating[i-1] {
			candyList[i] = candyList[i-1] + 1
		}
	}
	for i := num - 2; i >= 0; i-- {
		if rating[i] > rating[i+1] {
			if candyList[i] < candyList[i+1]+1 {
				candyList[i] = candyList[i+1] + 1
			}
		}
	}

	res := 0
	for i := 0; i < num; i++ {
		res += candyList[i]
	}
	return res + num

}