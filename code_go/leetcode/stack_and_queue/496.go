func nextGreaterElement01(nums []int) []int {
	res := make([]int, len(nums))
	stk := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stk) != 0 && stk[len(stk)-1] <= nums[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i] = -1
		} else {
			res[i] = stk[len(stk)-1]
		}
		stk = append(stk, nums[i])
	}
	return res
}