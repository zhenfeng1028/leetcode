package medium

// 关键：修改newInterval，并在恰当的时候加入答案
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	merged := false
	for _, interval := range intervals {
		if merged {
			ans = append(ans, interval)
			continue
		}
		// newInterval在当前区间的左侧且无交集
		if newInterval[1] < interval[0] {
			ans = append(ans, newInterval)
			ans = append(ans, interval)
			merged = true
		} else if newInterval[0] > interval[1] {
			// newInterval在当前区间的右侧且无交集
			ans = append(ans, interval)
		} else {
			// newInterval与当前区间一定存在交集，求它们的并集
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	if !merged {
		ans = append(ans, newInterval)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
