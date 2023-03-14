package medium

import "sort"

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		if idx == len(freq) || target < freq[idx][0] {
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数，若该数存在多个，可选多次，最大次数不能超过target
		most := min(target/freq[idx][0], freq[idx][1])
		for i := 1; i <= most; i++ {
			comb = append(comb, freq[idx][0])
			dfs(target-i*freq[idx][0], idx+1)
		}
		comb = comb[:len(comb)-most]
	}
	dfs(target, 0)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
