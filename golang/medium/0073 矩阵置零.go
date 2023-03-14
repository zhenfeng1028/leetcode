package medium

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowMap := make(map[int]struct{})
	colMap := make(map[int]struct{})
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rowMap[i] = struct{}{}
				colMap[j] = struct{}{}
			}
		}
	}
	for k := range rowMap {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}
	for k := range colMap {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}
