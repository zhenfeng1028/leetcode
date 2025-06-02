package medium

func subsets(nums []int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(index int)
	dfs = func(cur int) {
		if cur == len(nums) {
			comb := make([]int, len(tmp))
			copy(comb, tmp)
			res = append(res, comb)
			return
		}
		// 选择当前数
		tmp = append(tmp, nums[cur])
		dfs(cur + 1)
		tmp = tmp[:len(tmp)-1]
		// 不选当前数
		dfs(cur + 1)
	}
	dfs(0)
	return res
}
