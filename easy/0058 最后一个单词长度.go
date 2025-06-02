package easy

import "strings"

// 方法一：从后往前寻找不为空格的字符，并且遇到空格后终止
func lengthOfLastWord(s string) int {
	index1, index2 := -1, -1
	exist := false
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " && !exist {
			// 记录其下标
			index2 = i
			// 标记已存在非空格字符
			exist = true
		}
		if string(s[i]) == " " && exist {
			index1 = i
			break
		}
	}
	if index1 == -1 {
		return index2 + 1
	}
	return index2 - index1
}

// 方法二：使用golang自带的string split方法
func lengthOfLastWord2(s string) int {
	strs := strings.Split(s, " ")
	for i := len(strs) - 1; i >= 0; i-- {
		if len(strs[i]) != 0 {
			return len(strs[i])
		}
	}
	return 0
}
