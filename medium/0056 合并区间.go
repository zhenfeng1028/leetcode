package medium

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	merged := [][]int{}
	merged = append(merged, intervals[0])
	for i := 1; i < n; i++ {
		if intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
