package medium

// 使用辅助切片
func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, v := range matrix {
		for j, vv := range v {
			tmp[j][n-1-i] = vv
		}
	}
	copy(matrix, tmp)
}

// 原地旋转（题目要求）
func rotate2(matrix [][]int) {
	n := len(matrix)
	for row := 0; row < (n-1)/2; row++ {
		for col := 0; col < (n+1)/2; col++ {
			// 逆时针考虑，这时用一个tmp就可以解决问题
			tmp := matrix[row][col]
			matrix[row][col] = matrix[n-1-col][row]
			matrix[n-1-col][row] = matrix[n-1-row][n-1-col]
			matrix[n-1-row][n-1-col] = matrix[col][n-1-row]
			matrix[col][n-1-row] = tmp
		}
	}
}
