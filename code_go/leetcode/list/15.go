package list

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums) - 2; i++ {
		// 去重，需要第一个数不重复
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l, r := i + 1, len(nums) - 1
		for l < r {
			// 去重，也需要第一个数不重复
			if l > i + 1 && nums[l] == nums[l - 1] {
				l++
				continue
			}
			if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return res
}
