package medium

import "math"

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	curIdx := 0
	count := 0
	for curIdx < n {
		// 判断当前跳跃最大长度是否能达到最后一个元素
		if curIdx+nums[curIdx] >= n-1 {
			return count + 1
		} else {
			count++
			hashmap := make(map[int]int)
			// 若不能达到，记录可跳跃范围内所有元素能跳跃到的最大位置
			for i := curIdx + 1; i <= curIdx+nums[curIdx]; i++ {
				hashmap[i] = i + nums[i]
			}
			max := math.MinInt
			nextPos := -1
			for i, v := range hashmap {
				if v > max {
					max = v
					nextPos = i
				}
			}
			curIdx = nextPos
		}
	}
	return count
}
