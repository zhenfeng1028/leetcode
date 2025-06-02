package easy

import "sort"

func removeDuplicates(nums []int) int {
	hashmap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = struct{}{}
		}
	}
	s := make([]int, 0)
	for k := range hashmap {
		s = append(s, k)
	}
	sort.Ints(s)
	copy(nums, s)
	return len(s)
}
