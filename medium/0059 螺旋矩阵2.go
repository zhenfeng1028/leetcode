package medium

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	k := 1
	for k <= n*n {
		for i := left; i <= right; i++ {
			matrix[top][i] = k
			k++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = k
			k++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = k
			k++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = k
			k++
		}
		left++
	}
	return matrix
}
