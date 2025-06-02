package easy

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left != right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if target > nums[left] {
		return left + 1
	}
	return left
}
