package medium

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	var dfs func(x, y int, cur int) bool
	dfs = func(x, y int, cur int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[cur] || used[x][y] || cur == len(word) {
			return false
		}
		used[x][y] = true
		if cur == len(word)-1 {
			return true
		}
		if dfs(x-1, y, cur+1) {
			return true
		}
		if dfs(x+1, y, cur+1) {
			return true
		}
		if dfs(x, y-1, cur+1) {
			return true
		}
		if dfs(x, y+1, cur+1) {
			return true
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
