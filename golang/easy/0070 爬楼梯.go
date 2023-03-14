package easy

// 递归
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 动态规划
func climbStairs2(n int) int {
	dp := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		// 最后一步可能跨了一级台阶，也可能跨了两级台阶
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
