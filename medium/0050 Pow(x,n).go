package medium

// 暴力解法 时间复杂度：O(n)
func myPow(x float64, n int) float64 {
	ans := 1.0
	sign := -1
	if n < 0 {
		n = -n
	} else {
		sign = 1
	}
	for i := 0; i < n; i++ {
		ans = ans * x
	}
	if sign == 1 {
		return ans
	} else {
		return 1 / ans
	}
}

// 分治 时间复杂度：O(logn)
func myPow2(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n > 0 {
		return myPowHelper(x, n)
	} else {
		return 1 / myPowHelper(x, n)
	}
}

func myPowHelper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 1 {
		half := myPowHelper(x, n/2)
		return half * half * x
	} else {
		half := myPowHelper(x, n/2)
		return half * half
	}
}
