func nextGreaterElement03(nums []int) []int {
	res := make([]int, len(nums))
	stk := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stk) != 0 && nums[stk[len(stk)-1]] <= nums[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i] = 0
		} else {
			res[i] = stk[len(stk)-1] - i
		}
		stk = append(stk, i)
	}
	return res
}