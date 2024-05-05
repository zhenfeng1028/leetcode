package medium

/*
	给你一个大小为 m x n 的二进制矩阵 grid 。

	岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

	岛屿的面积是岛上值为 1 的单元格的数目。

	计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
*/

// 方法一：深度优先搜索
func dfs(grid [][]int, cur_i, cur_j int) int {
	if cur_i < 0 || cur_j < 0 || cur_i == len(grid) || cur_j == len(grid[0]) || grid[cur_i][cur_j] != 1 {
		return 0
	}
	grid[cur_i][cur_j] = 0
	di := [4]int{-1, 1, 0, 0}
	dj := [4]int{0, 0, -1, 1}
	ans := 1
	for index := 0; index < 4; index++ {
		next_i, next_j := cur_i+di[index], cur_j+dj[index]
		ans += dfs(grid, next_i, next_j)
	}
	return ans
}

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = max(ans, dfs(grid, i, j))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法二：深度优先搜索+栈
func maxAreaOfIsland2(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cur := 0
			stacki := []int{i}
			stackj := []int{j}
			for len(stacki) != 0 {
				cur_i, cur_j := stacki[len(stacki)-1], stackj[len(stackj)-1]
				if len(stacki) > 0 {
					stacki = stacki[:len(stacki)-1]
				}
				if len(stackj) > 0 {
					stackj = stackj[:len(stackj)-1]
				}
				if cur_i < 0 || cur_j < 0 || cur_i == len(grid) || cur_j == len(grid[0]) || grid[cur_i][cur_j] != 1 {
					continue
				}
				cur++
				grid[cur_i][cur_j] = 0
				di := [4]int{-1, 1, 0, 0}
				dj := [4]int{0, 0, -1, 1}
				for index := 0; index < 4; index++ {
					next_i, next_j := cur_i+di[index], cur_j+dj[index]
					stacki = append(stacki, next_i)
					stackj = append(stackj, next_j)
				}
			}
			ans = max(ans, cur)
		}
	}
	return ans
}

// 方法三：广度优先所有（将方法二中的栈改为队列）
func maxAreaOfIsland3(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cur := 0
			stacki := []int{i}
			stackj := []int{j}
			for len(stacki) != 0 {
				cur_i, cur_j := stacki[0], stackj[0]
				if len(stacki) > 0 {
					stacki = stacki[1:]
				}
				if len(stackj) > 0 {
					stackj = stackj[1:]
				}
				if cur_i < 0 || cur_j < 0 || cur_i == len(grid) || cur_j == len(grid[0]) || grid[cur_i][cur_j] != 1 {
					continue
				}
				cur++
				grid[cur_i][cur_j] = 0
				di := [4]int{-1, 1, 0, 0}
				dj := [4]int{0, 0, -1, 1}
				for index := 0; index < 4; index++ {
					next_i, next_j := cur_i+di[index], cur_j+dj[index]
					stacki = append(stacki, next_i)
					stackj = append(stackj, next_j)
				}
			}
			ans = max(ans, cur)
		}
	}
	return ans
}
