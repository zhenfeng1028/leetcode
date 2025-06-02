package medium

// 不包含重复数字
func permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 1 {
		return append(ans, nums)
	}
	for i := 0; i < len(nums); i++ {
		tmp := []int{}
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		ret := permute(tmp)
		for _, v := range ret {
			tmp2 := []int{}
			tmp2 = append(tmp2, nums[i])
			tmp2 = append(tmp2, v...)
			ans = append(ans, tmp2)
		}
	}
	return ans
}
