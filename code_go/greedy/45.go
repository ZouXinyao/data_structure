package main

func jump(nums []int) int {
	length := len(nums)
	// end相当于1个临时变量，保存当前这一步能走多远，当走到这么远的时候，就需要走下一步了
	end := 0
	maxPosition := 0
	steps := 0
	// 这里length - 1是因为：
	// 1、i == end时走1步，当在倒数第二个位置时，刚好i == end，那么step++走1步保证能到最后;
	// 2、i != end，那么就更不需要判断最后一位了，因为一定可以达到
	for i := 0; i < length - 1; i++ {
		maxPosition = max(maxPosition, i + nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y { return x }
	return y
}

//func main() {
//	nums := []int{1}
//	fmt.Println(jump(nums))
//}
