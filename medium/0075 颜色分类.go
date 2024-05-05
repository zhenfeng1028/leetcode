package medium

func sortColors(nums []int) {
	for i := len(nums); i > 1; i-- {
		for j := 1; j < i; j++ {
			if nums[j-1] > nums[j] {
				swap(nums, j-1, j)
			}
		}
	}
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
