package medium

import (
	"reflect"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	tmp := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for m := j + 1; m < len(nums)-1; m++ {
				for n := m + 1; n < len(nums); n++ {
					if nums[i]+nums[j]+nums[m]+nums[n] == target {
						tuple := make([]int, 0)
						tuple = append(tuple, nums[i])
						tuple = append(tuple, nums[j])
						tuple = append(tuple, nums[m])
						tuple = append(tuple, nums[n])
						tmp = append(tmp, tuple)
					}
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
