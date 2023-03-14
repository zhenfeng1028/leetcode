package medium

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	subStr := make([]map[string]struct{}, 0)
	for i := 0; i < len(s); i++ {
		hashmap := make(map[string]struct{})
		hashmap[string(s[i])] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			if _, ok := hashmap[string(s[j])]; !ok {
				hashmap[string(s[j])] = struct{}{}
			} else {
				break
			}
		}
		subStr = append(subStr, hashmap)
	}
	maxLen := math.MinInt
	for _, v := range subStr {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	return maxLen
}
