package medium

// 用dfs会超时
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	var cnt int
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x >= m || y >= n || obstacleGrid[x][y] == 1 {
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

// 动态规划
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		if obstacleGrid[i-1][0] == 0 {
			dp[i][1] = 1
		} else {
			break // 遇到第一个障碍物，后面的都是不可达的
		}
	}
	for i := 1; i <= n; i++ {
		if obstacleGrid[0][i-1] == 0 {
			dp[1][i] = 1
		} else {
			break // 遇到第一个障碍物，后面的都是不可达的
		}
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			if obstacleGrid[i-1][j-1] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}
