package medium

import "sort"

// 1.从后向前寻找第一个顺序对(i,i+1)，满足a[i] < a[i+1]
// 2.如果找到了，在区间[i+1,n)中从后向前寻找第一个元素j，满足a[j]>a[i]
// 3.交换a[i]与a[j]，并对区间[i+1,n)进行升序排序
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			index = i - 1
			break
		}
	}
	if index == -1 {
		// 不存在下一个更大的排列
		sort.Ints(nums)
		return
	}
	for j := len(nums) - 1; j > index; j-- {
		if nums[j] > nums[index] {
			tmp := nums[index]
			nums[index] = nums[j]
			nums[j] = tmp
			// 将下标index+1到最后的元素按从小到大的顺序进行排序
			sort.Ints(nums[index+1:])
			return
		}
	}
}
