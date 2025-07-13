package easy

import "strings"

/*
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
*/

func isPalindrome(s string) bool {
	alpNumStr := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			alpNumStr += string(c + 32)
		} else if c >= 'a' && c <= 'z' {
			alpNumStr += string(c)
		} else if c >= '0' && c <= '9' {
			alpNumStr += string(c)
		}
	}
	left, right := 0, len(alpNumStr)-1
	for left < right {
		if alpNumStr[left] != alpNumStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isLetterOrNum(s[left]) {
			left++
		}
		for left < right && !isLetterOrNum(s[right]) {
			right--
		}
		leftChar := strings.ToLower(string(s[left]))
		rightChar := strings.ToLower(string(s[right]))
		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetterOrNum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
