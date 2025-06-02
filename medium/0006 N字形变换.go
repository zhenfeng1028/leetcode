package medium

import "math"

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	// 分组，确定每组字符个数
	groupCharNum := numRows*2 - 2
	// 确定列数
	cols := int(math.Ceil(float64(len(s)) / float64(groupCharNum) * float64(numRows-1)))
	// 构建一个二维数组
	arr2d := [][]string{}
	// 填充字符
	ptr := 0
	for j := 0; j < cols; j++ {
		arr := []string{}
		for i := 0; i < numRows; i++ {
			if j%(numRows-1) == 0 {
				if ptr >= len(s) {
					arr = append(arr, " ")
					continue
				}
				arr = append(arr, string(s[ptr]))
				ptr++
			} else {
				if i == numRows-(j%(numRows-1))-1 {
					if ptr >= len(s) {
						arr = append(arr, " ")
						continue
					}
					arr = append(arr, string(s[ptr]))
					ptr++
				} else {
					arr = append(arr, " ")
				}
			}
		}
		arr2d = append(arr2d, arr)
	}

	ans := ""
	rowLen := len(arr2d)
	colLen := numRows
	for j := 0; j < colLen; j++ {
		for i := 0; i < rowLen; i++ {
			if arr2d[i][j] != " " {
				ans = ans + arr2d[i][j]
			}
		}
	}
	return ans
}

// 寻找转换后下标的规律
func convert2(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}
