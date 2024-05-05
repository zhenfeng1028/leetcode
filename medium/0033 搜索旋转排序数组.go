package medium

// 先找出旋转位置的下标，再使用二分法
// 时间复杂度：O(n+logn) => O(n)
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 寻找旋转位置的下标
	i := 0
	for i < len(nums)-1 {
		if nums[i] < nums[i+1] {
			i++
		} else {
			break
		}
	}
	left, right := -1, -1
	if nums[0] <= target {
		left = 0
		right = i
	}
	if nums[0] > target {
		left = i + 1
		right = len(nums) - 1
	}
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// 直接使用二分法
// 时间复杂度：O(logn)
func search2(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 左半部分有序
		if nums[0] <= nums[mid] {
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
	return -1
}
