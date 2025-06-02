package medium

// 按层模拟
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var (
		rows, cols               = len(matrix), len(matrix[0])
		order                    = make([]int, rows*cols)
		index                    = 0
		left, right, top, bottom = 0, cols - 1, 0, rows - 1
	)
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			order[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				order[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}
