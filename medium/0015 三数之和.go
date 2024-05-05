package medium

import (
	"reflect"
	"sort"
)

// 方法一：暴力解法（找出所有满足条件的组合，再进行去重）
func threeSum(nums []int) [][]int {
	tmp := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tuple := make([]int, 0)
					tuple = append(tuple, nums[i])
					tuple = append(tuple, nums[j])
					tuple = append(tuple, nums[k])
					tmp = append(tmp, tuple)
				}
			}
		}
	}
	ans := [][]int{}
	if len(tmp) > 0 {
		ans = append(ans, tmp[0])
	}
	for i := 1; i < len(tmp); i++ {
		found := false
		for j := 0; j < len(ans); j++ {
			if equalSlice(ans[j], tmp[i]) {
				found = true
				break
			}
		}
		if !found {
			ans = append(ans, tmp[i])
		}
	}
	return ans
}

func equalSlice(s1, s2 []int) bool {
	sort.Ints(s1)
	sort.Ints(s2)
	return reflect.DeepEqual(s1, s2)
}

// 方法二：排序+双指针
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	// 枚举a
	for i := 0; i < len(nums); i++ {
		// 需要和上一次枚举的数不同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		// c对应的指针初始指向数组的最右端
		k := len(nums) - 1
		// 枚举b
		for j := i + 1; j < len(nums); j++ {
			// 需要和上一次枚举的数不同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 需要保证b的指针在c的指针的左侧
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			// 如果指针重合，随着b后续的增加
			// 不会再有满足a+b+c=0的c了，可以退出循环
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				s := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, s)
			}
		}
	}
	return ans
}
