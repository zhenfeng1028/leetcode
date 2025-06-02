package easy

// 方法一：暴力枚举
func twoSum(nums []int, target int) []int {
	indices := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices = append(indices, i)
				indices = append(indices, j)
				return indices
			}
		}
	}

	return nil
}

// 方法二：哈希表
func twoSum2(nums []int, target int) []int {
	numsmap := make(map[int]int)
	indices := []int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := numsmap[target-nums[i]]; ok {
			indices = append(indices, i)
			indices = append(indices, j)
			return indices
		}
		numsmap[nums[i]] = i
	}

	return nil
}
