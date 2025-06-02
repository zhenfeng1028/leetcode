package medium

import (
	"math"
	"strconv"
)

// 1.获取第一个有效字符（-、+或者数字）
// 2.获取第一个无效字符（空格及其他非数字字符）或不存在
// 3.判断其是否小于-2^31或者大于2^31
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	signIdx, firstNumIdx, lastNumIdx := -1, -1, -1
	for i := 0; i < len(s); i++ {
		if int(byte(s[i])) > 48 && int(byte(s[i])) <= 57 && firstNumIdx == -1 {
			firstNumIdx = i
		}
		if (string(s[i]) == "-" || string(s[i]) == "+") && signIdx == -1 && firstNumIdx == -1 {
			if i+1 < len(s) && int(byte(s[i+1])) > 48 && int(byte(s[i+1])) <= 57 {
				signIdx = i
				firstNumIdx = i + 1
				i++
			}
		}
		if (int(byte(s[i])) < 48 || int(byte(s[i])) > 57) && firstNumIdx != -1 {
			lastNumIdx = i - 1
			break
		}
		if signIdx == -1 && firstNumIdx == -1 && s[i] != byte(' ') && s[i] != byte('0') {
			return 0
		}
	}
	if lastNumIdx == -1 {
		lastNumIdx = len(s) - 1
	}
	numStr := ""
	// 未找到有效数字
	if firstNumIdx == -1 {
		return 0
	}
	numStr = s[firstNumIdx : lastNumIdx+1]
	maxNum := strconv.Itoa(int(math.MaxInt32))
	maxLen := len(maxNum)
	num, _ := strconv.Atoi(numStr)
	// 未找到有效符号或者为+
	if signIdx == -1 || (signIdx != -1 && string(s[signIdx]) == "+") {
		if len(numStr) < maxLen {
			return num
		}
		if len(numStr) > maxLen {
			return math.MaxInt32
		}
		// 长度相等从左到右依次比较每个位
		for i, c := range numStr {
			if maxNum[i] < byte(c) {
				return math.MaxInt32
			}
		}
		return num
	} else {
		if len(numStr) < maxLen {
			return -num
		}
		if len(numStr) > maxLen {
			return math.MinInt32
		}
		// 长度相等从左到右依次比较每个位
		for i, c := range numStr {
			if maxNum[i] < byte(c) {
				return math.MinInt32
			}
		}
		return -num
	}
}
