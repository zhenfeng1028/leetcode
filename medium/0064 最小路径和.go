package medium

import "math"

func minPathSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	minSum := math.MaxInt64
	var dfs func(x, y int, cur int)
	dfs = func(x, y int, cur int) {
		if cur >= minSum { // 剪枝
			return
		}
		if x >= r || y >= c {
			return
		}
		cur += grid[x][y]
		if x == r-1 && y == c-1 {
			minSum = min(minSum, cur)
		}
		dfs(x, y+1, cur)
		dfs(x+1, y, cur)
	}
	dfs(0, 0, 0)
	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 动态规划
func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][1] = dp[i-1][1] + grid[i-1][0]
	}
	for j := 1; j <= n; j++ {
		dp[1][j] = dp[1][j-1] + grid[0][j-1]
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[m][n]
}
