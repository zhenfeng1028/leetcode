package medium

// 双指针 时间复杂度：O(n)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < target {
			left++
		}
		if nums[right] > target {
			right--
		}
		if nums[left] == target && nums[right] == target {
			break
		}
	}
	if left < right {
		return []int{left, right}
	} else if left == right && nums[left] == target {
		return []int{left, right}
	} else {
		return []int{-1, -1}
	}
}

// 二分法 时间复杂度：O(logn)
func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	// 寻找第一个位置的下标
	index1 := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			index1 = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 寻找第二个位置的下标
	left, right = 0, len(nums)-1
	index2 := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			index2 = mid
			left = mid + 1
		}
	}
	if index1 > index2 || index1 == -1 {
		return []int{-1, -1}
	}
	return []int{index1, index2}
}
