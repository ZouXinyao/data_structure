func nextGreaterElement02(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stk := []int{}
	for i := 2*n - 1; i >= 0; i-- {
		for len(stk) != 0 && stk[len(stk)-1] <= nums[i%n] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = stk[len(stk)-1]
		}
		stk = append(stk, nums[i%n])
	}
	return res
}