package medium

func combine(n int, k int) [][]int {
	res := [][]int{}
	tmp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝
		if len(tmp)+(n-cur+1) < k {
			return
		}
		if len(tmp) == k {
			comb := make([]int, k)
			copy(comb, tmp)
			res = append(res, comb)
			return
		}
		// 加了剪枝条件后，该条件将不可达，可以省去
		// if cur == n+1 {
		// 	return
		// }
		// 考虑选择当前位置
		tmp = append(tmp, cur)
		dfs(cur + 1)
		tmp = tmp[:len(tmp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return res
}
