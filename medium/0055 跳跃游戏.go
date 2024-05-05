package medium

func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxIndex {
			maxIndex = max(maxIndex, i+nums[i])
		}
		if maxIndex >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
