package dp

// 回溯法
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	half := sum / 2
	if half * 2 != sum { return false }

	tempSum := 0
	var helper func(idx int) bool
	helper = func(idx int) bool {
		if tempSum == half {
			return true
		}

		for i := idx; i < len(nums); i++ {
			if tempSum + nums[i] > half { continue }
			tempSum += nums[i]
			if helper(i + 1) { return true }
			tempSum -= nums[i]
		}
		return false
	}
	return helper(0)
}

// 史上最low的DP，有一些优化，比如最大值>half的话直接返回false；可以使用1维数组的DP
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	half := sum / 2
	if half * 2 != sum { return false }

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			temp := make([]int, half + 1)
			for j := 1; j < half + 1; j++ {
				if j >= nums[0] { temp[j] = nums[0] }
			}
			dp[i] = temp
			continue
		}
		dp[i] = make([]int, half + 1)
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j < half + 1; j++ {
			if j >= nums[i] {
				dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - nums[i]] + nums[i])
			} else {
				dp[i][j] = dp[i - 1][j]
			}
		}
	}

	return dp[len(nums) - 1][half] == half
}

func max(a, b int) int {
	if a > b { return a }
	return b
}