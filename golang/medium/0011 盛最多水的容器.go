package medium

import "math"

// 暴力解法
func maxArea(height []int) int {
	maxArea := math.MinInt
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 双指针
func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	maxArea := math.MinInt
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
