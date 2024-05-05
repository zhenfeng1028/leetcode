package easy

func plusOne(digits []int) []int {
	if len(digits) == 0 && digits[0] == 0 {
		return []int{1}
	}
	digits[len(digits)-1] += 1
	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			digits[i] += 1
		}
		if digits[i] == 10 {
			digits[i] = 0
			carry = true
		} else {
			break
		}
	}
	if digits[0] != 0 {
		return digits
	} else {
		ans := []int{1}
		ans = append(ans, digits...)
		return ans
	}
}
