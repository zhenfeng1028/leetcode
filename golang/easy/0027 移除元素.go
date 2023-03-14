package easy

func removeElement(nums []int, val int) int {
	hashmap := make(map[int]int)
	for _, num := range nums {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = 1
		} else {
			hashmap[num]++
		}
	}
	delete(hashmap, val)
	s := make([]int, 0)
	for k, count := range hashmap {
		for i := 0; i < count; i++ {
			s = append(s, k)
		}
	}
	copy(nums, s)
	return len(s)
}
