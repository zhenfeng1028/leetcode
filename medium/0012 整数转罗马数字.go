package medium

func intToRoman(num int) string {
	ans := ""
	thousandDigit := num / 1000
	for thousandDigit > 0 {
		ans = ans + "M"
		thousandDigit--
	}
	hundredDigit := num / 100 % 10
	if hundredDigit > 0 && hundredDigit < 4 {
		for hundredDigit > 0 {
			ans = ans + "C"
			hundredDigit--
		}
	} else if hundredDigit == 4 {
		ans = ans + "CD"
	} else if hundredDigit == 5 {
		ans = ans + "D"
	} else if hundredDigit > 5 && hundredDigit < 9 {
		ans = ans + "D"
		for hundredDigit > 5 {
			ans = ans + "C"
			hundredDigit--
		}
	} else if hundredDigit == 9 {
		ans = ans + "CM"
	}
	tensDigit := num % 100 / 10
	if tensDigit > 0 && tensDigit < 4 {
		for tensDigit > 0 {
			ans = ans + "X"
			tensDigit--
		}
	} else if tensDigit == 4 {
		ans = ans + "XL"
	} else if tensDigit == 5 {
		ans = ans + "L"
	} else if tensDigit > 5 && tensDigit < 9 {
		ans = ans + "L"
		for tensDigit > 5 {
			ans = ans + "X"
			tensDigit--
		}
	} else if tensDigit == 9 {
		ans = ans + "XC"
	}
	singleDigit := num % 10
	if singleDigit > 0 && singleDigit < 4 {
		for singleDigit > 0 {
			ans = ans + "I"
			singleDigit--
		}
	} else if singleDigit == 4 {
		ans = ans + "IV"
	} else if singleDigit == 5 {
		ans = ans + "V"
	} else if singleDigit > 5 && singleDigit < 9 {
		ans = ans + "V"
		for singleDigit > 5 {
			ans = ans + "I"
			singleDigit--
		}
	} else if singleDigit == 9 {
		ans = ans + "IX"
	}
	return ans
}
