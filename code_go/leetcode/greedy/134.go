package main

// total用来计算总差值，<0代表油不够
// tempSum<0时，代表，绝对不能让i前面的位置作起点，因为如果做了起点，这段局部距离的油就不够了。
func canCompleteCircuit(gas []int, cost []int) int {
	tempSum, totalSum := 0, 0
	start := 0
	for i:=0; i<len(gas); i++ {
		tempSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if tempSum < 0 {
			start = i + 1
			tempSum = 0
		}
	}
	if totalSum < 0 { return -1 }
	return start
}
