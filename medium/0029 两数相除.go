package medium

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		} else if divisor == -1 {
			return math.MaxInt32
		}
	}
	if dividend == math.MaxInt32 {
		if divisor == 1 {
			return math.MaxInt32
		} else if divisor == -1 {
			return math.MinInt32 + 1
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}
	sign := ""
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		sign = "-"
	}
	if dividend < 0 && divisor > 0 {
		dividend = -dividend
		sign = "-"
	}
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}
	sum := divisor
	count := 0
	for sum <= dividend {
		count++
		sum += divisor
	}
	if sign == "-" {
		return -count
	} else {
		return count
	}
}
