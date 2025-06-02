package medium

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		if _, ok := hashmap[sortedStr]; !ok {
			hashmap[sortedStr] = make([]string, 0)
		}
		hashmap[sortedStr] = append(hashmap[sortedStr], str)
	}
	ans := [][]string{}
	for _, group := range hashmap {
		ans = append(ans, group)
	}
	return ans
}
