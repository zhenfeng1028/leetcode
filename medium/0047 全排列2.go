package medium

// 可能包含重复数字
func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 1 {
		return append(ans, nums)
	}
	hashmap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := hashmap[nums[i]]; ok {
			continue
		} else {
			hashmap[nums[i]] = struct{}{}
		}
		tmp := []int{}
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		ret := permuteUnique(tmp)
		for _, v := range ret {
			tmp2 := []int{}
			tmp2 = append(tmp2, nums[i])
			tmp2 = append(tmp2, v...)
			ans = append(ans, tmp2)
		}
	}
	return ans
}
