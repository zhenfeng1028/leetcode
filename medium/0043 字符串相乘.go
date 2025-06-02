package medium

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := []string{}
	for i := len(num1) - 1; i >= 0; i-- {
		carry := 0
		ss := []string{}
		for j := len(num2) - 1; j >= 0; j-- {
			result := int(num1[i]-48) * int(num2[j]-48)
			if carry > 0 {
				result += carry
			}
			remainder := result % 10
			carry = result / 10
			if j == 0 {
				ss = append(ss, strconv.Itoa(result))
			} else {
				ss = append(ss, strconv.Itoa(remainder))
			}

		}
		str := ""
		for k := len(ss) - 1; k >= 0; k-- {
			str += ss[k]
		}
		tmp = append(tmp, str)
	}
	first := tmp[0]
	for k := 1; k < len(tmp); k++ {
		second := tmp[k]
		first = add(first, second, k)
	}
	return first
}

func add(first, second string, tail int) string {
	ss := []string{}
	ss = append(ss, string(first[len(first)-tail:]))
	first = first[:len(first)-tail]
	n1 := len(first)
	n2 := len(second)
	carry := 0
	for n1-1 >= 0 && n2-1 >= 0 {
		sum := int(first[n1-1]-48) + int(second[n2-1]-48) + carry
		remainder := sum % 10
		carry = sum / 10
		if n1-1 > 0 || n2-1 > 0 {
			ss = append(ss, strconv.Itoa(remainder))
		}
		if n1-1 == 0 && n2-1 == 0 {
			ss = append(ss, strconv.Itoa(sum))
			break
		}
		n1--
		n2--
	}
	if n1 <= 0 {
		second = second[:n2]
		if carry == 0 {
			ss = append(ss, second)
		} else {
			for k := len(second) - 1; k >= 0; k-- {
				sum := int(second[k]-48) + carry
				remainder := sum % 10
				carry = sum / 10
				ss = append(ss, strconv.Itoa(remainder))
			}
			if carry > 0 {
				ss = append(ss, strconv.Itoa(carry))
			}
		}
	}
	if n2 <= 0 {
		first = first[:n1]
		if carry == 0 {
			ss = append(ss, first)
		} else {
			for k := len(first) - 1; k >= 0; k-- {
				sum := int(first[k]-48) + carry
				remainder := sum % 10
				carry = sum / 10
				ss = append(ss, strconv.Itoa(remainder))
			}
			if carry > 0 {
				ss = append(ss, strconv.Itoa(carry))
			}
		}
	}
	str := ""
	for k := len(ss) - 1; k >= 0; k-- {
		str += ss[k]
	}
	return str
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
