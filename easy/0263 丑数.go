package easy

/*
	丑数 就是只包含质因数 2、3 和 5 的正整数。
*/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	factors := []int{2, 3, 5}
	for _, factor := range factors {
		for n%factor == 0 {
			n = n / factor
		}
	}
	return n == 1
}
