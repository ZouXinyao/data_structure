package list

func moveZeroes(nums []int)  {
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[l] == 0 && nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		} else if nums[l] != 0 {
			l++
		}
	}
}
