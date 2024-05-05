package medium

import (
	"math"
)

// 方法一：暴力解法
// 1.找出所有子串
// 2.依次判断是否是回文子串，如果是就记录下来
func longestPalindrome(s string) string {
	allSubStr := []string{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			allSubStr = append(allSubStr, s[i:j])
		}
	}
	maxLen := math.MinInt
	longestPalStr := ""
	for _, substr := range allSubStr {
		if isPalindrome(substr) {
			if len(substr) > maxLen {
				maxLen = len(substr)
				longestPalStr = substr
			}
		}
	}
	return longestPalStr
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 方法二：中心扩散法
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := math.MinInt
	// 记录最长回文子串的起始下标
	begin := 0
	for i := 0; i < len(s)-1; i++ {
		oddLen := expandAroundCenter(s, i, i)
		evenLen := expandAroundCenter(s, i, i+1)
		curLen := max(oddLen, evenLen)
		if curLen > maxLen {
			maxLen = curLen
			begin = i - (maxLen-1)/2
		}
	}
	return s[begin : begin+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法三：动态规划
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := math.MinInt
	// 记录最长回文子串的起始下标
	begin := 0
	// dp[i][j]表示s[i...j]是否是回文
	dp := [1000][1000]bool{}
	// 单个字符一定是回文，将表格对角线设置为true
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	// 注意：只需要填写表格的上半部份，从上往下，从左往右填写
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 只要dp[i][j] = true就表示s[i...j]是回文
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
