package medium

func isValidSudoku(board [][]byte) bool {

	// 遍历所有行
	for _, row := range board {
		hashmap := make(map[string]int)
		for _, c := range row {
			hashmap[string(c)]++
		}
		for _, v := range hashmap {
			if v > 1 {
				return false
			}
		}
	}
	// 遍历所有列
	for i := 0; i < 9; i++ {
		hashmap := make(map[string]int)
		for _, row := range board {
			hashmap[string(row[i])]++
		}
		for _, v := range hashmap {
			if v > 1 {
				return false
			}
		}
	}
	// 遍历所有九宫格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			hashmap := make(map[string]int)
			for k, row := range board {
				if k >= 3*i && k < 3*(i+1) {
					for l, c := range row {
						if l >= 3*j && l < 3*(j+1) {
							hashmap[string(c)]++
						}
					}
				}
			}
			for _, v := range hashmap {
				if v > 1 {
					return false
				}
			}
		}
	}
	return true
}
