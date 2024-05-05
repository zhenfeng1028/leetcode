package hard

// 动态规划：d[i][j]表示字符串s的前i个字符和模式p的前j个字符是否能匹配
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true // 当字符串s和模式p均为空时，匹配成功
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true // 因为*才能匹配空字符串，所以只有模式p的前j个字符均是*时，dp[0][j]才为真
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' { // 如果p[j-1]是*，对s[i-1]没有任何要求，但是*可以匹配零或任意多个小写字母，因此状态转移方程分两种情况，即使用或不使用这个*
				dp[i][j] = dp[i][j-1] || dp[i-1][j] // 如果不使用*，那么就会从dp[i][j-1]转移而来；如果使用*，那么就会从dp[i-1][j]转移而来
			}
		}
	}
	return dp[m][n]
}
