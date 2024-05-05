package medium

import "math"

// 暴力解法 时间复杂度：O(n^2)
func maxSubArray(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// 动态规划 时间复杂度：O(n)
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxNum := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxNum {
			maxNum = dp[i]
		}
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
