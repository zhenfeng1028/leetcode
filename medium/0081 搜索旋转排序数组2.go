package medium

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target || nums[right] == target {
			return true
		}
		left++
		right--
	}
	return false
}

// 官方解法：二分扩展
func search2(nums []int, target int) bool {
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		// 无法判断哪个区间有序，将左边界加一，右边界减一
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] { // 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
