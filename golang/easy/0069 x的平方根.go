package easy

func mySqrt(x int) int {
	ans := -1
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid <= x {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}
