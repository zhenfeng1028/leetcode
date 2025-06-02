package medium

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	cnt := 1
	i := 1
	begin, end := 0, 0 // 多余元素的起始下标和结束下标
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			if cnt == 2 {
				begin = i
			}
			cnt++
		} else {
			if cnt > 2 {
				end = i
				nums = append(nums[:begin], nums[end:]...)
				i -= end - begin
			}
			cnt = 1
		}
		i++
	}
	return len(nums)
}
