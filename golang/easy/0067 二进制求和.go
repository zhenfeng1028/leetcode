package easy

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")
	a2, b2 := []int{}, []int{}
	for _, v := range a1 {
		v1, _ := strconv.Atoi(v)
		a2 = append(a2, v1)
	}
	for _, v := range b1 {
		v1, _ := strconv.Atoi(v)
		b2 = append(b2, v1)
	}
	diff := len(a2) - len(b2)
	zeros := []int{}
	if diff > 0 {
		for diff > 0 {
			zeros = append(zeros, 0)
			diff--
		}
		zeros = append(zeros, b2...)
		b2 = zeros
	} else if diff < 0 {
		for diff < 0 {
			zeros = append(zeros, 0)
			diff++
		}
		zeros = append(zeros, a2...)
		a2 = zeros
	}
	carry := false
	tmp := []int{}
	for i := len(a2) - 1; i >= 0; i-- {
		tmpV := a2[i] + b2[i]
		if carry {
			tmpV += 1
		}
		if tmpV == 2 {
			tmp = append(tmp, 0)
			carry = true
		} else if tmpV == 3 {
			tmp = append(tmp, 1)
			carry = true
		} else {
			tmp = append(tmp, a2[i]+b2[i])
			carry = false
		}
	}
	tmpStr := []string{}
	for _, v := range tmp {
		tmpStr = append(tmpStr, strconv.Itoa(v))
	}
	return strings.Join(tmpStr, "")
}
