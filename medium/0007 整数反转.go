package medium

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	tmp := 0
	for x > 0 {
		tmp = tmp*10 + x%10
		x /= 10
	}
	if tmp > math.MaxInt32 {
		return 0
	}
	return tmp * sign
}
