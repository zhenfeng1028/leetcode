package medium

// 用dfs会超时
func uniquePaths(m int, n int) int {
	var cnt int
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x >= m || y >= n {
			return
		}
		if x == m-1 && y == n-1 {
			cnt++
			return
		}
		dfs(x, y+1) // 向右
		dfs(x+1, y) // 向下
	}
	dfs(0, 0)
	return cnt
}

func uniquePaths2(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
