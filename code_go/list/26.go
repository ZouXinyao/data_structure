package list

// 双指针，只用管前面不重复的元素，后面的那些不用管。
func removeDuplicates(nums []int) int {
	if len(nums) < 2 { return len(nums) }
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}
	return l+1
}
