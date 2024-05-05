package medium

import "math"

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxFloat64
	closest := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if math.Abs(float64(sum-target)) < diff {
					diff = math.Abs(float64(sum - target))
					closest = sum
				}
			}
		}
	}
	return closest
}
