package medium

// 动态规划：dp[i]表示字符串s的前i个字符的解码方法数
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		cnt1, cnt2 := 0, 0 // 保存两种情况的解码方法数
		// 情况一：使用一个字符
		if s[i-1] != '0' {
			cnt1 = dp[i-1]
		}
		// 情况二：使用两个字符
		if i > 1 && s[i-2] != '0' && 10*(s[i-2]-48)+(s[i-1]-48) <= 26 {
			cnt2 = dp[i-2]
		}
		dp[i] = cnt1 + cnt2
	}
	return dp[n]
}
