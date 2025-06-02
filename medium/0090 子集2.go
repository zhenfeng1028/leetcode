package medium

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	tmp := []int{}
	sort.Ints(nums)
	var dfs func(choosePre bool, cur int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			comb := make([]int, len(tmp))
			copy(comb, tmp)
			res = append(res, comb)
			return
		}
		// 不选当前数
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur] == nums[cur-1] {
			return
		}
		// 选择当前数
		tmp = append(tmp, nums[cur])
		dfs(true, cur+1)
		tmp = tmp[:len(tmp)-1]
	}
	dfs(false, 0)
	return res
}
