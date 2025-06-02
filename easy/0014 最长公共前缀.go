package easy

import "math"

// 纵向扫描
func longestCommonPrefix(strs []string) string {
	// 如果数组只含有一个字符串
	if len(strs) == 1 {
		return strs[0]
	}
	// 找出数组中长度最短的字符串
	minIndex := 0
	minStrLen := math.MaxInt
	for i, str := range strs {
		if len(str) < minStrLen {
			minIndex = i
			minStrLen = len(str)
		}
	}
	if minStrLen == 0 {
		return ""
	}
	minStr := strs[minIndex]
	for i, c := range minStr {
		for _, str := range strs {
			if string(c) != string(str[i]) {
				if i == 0 {
					return ""
				} else {
					return minStr[0:i]
				}
			}
		}
	}
	return minStr
}

// 横向扫描
func longestCommonPrefix2(strs []string) string {
	// 如果数组只含有一个字符串
	if len(strs) == 1 {
		return strs[0]
	}
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		str = LongestCommonPrefix(str, strs[i])
	}
	return str
}

func LongestCommonPrefix(str1, str2 string) string {
	minLen := 0
	minStr := ""
	if len(str1) < len(str2) {
		minLen = len(str1)
		minStr = str1
	} else {
		minLen = len(str2)
		minStr = str2
	}
	for i := 0; i < minLen; i++ {
		if string(str1[i]) != string(str2[i]) {
			if i == 0 {
				return ""
			} else {
				return str1[0:i]
			}
		}
	}
	return minStr
}
