package medium

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	object := countAndSay(n - 1)
	count := 1
	tmp := []string{}
	for i := 0; i < len(object); i++ {
		if i != len(object)-1 && object[i+1] == object[i] {
			count++
			if i == len(object)-2 {
				tmp = append(tmp, strconv.Itoa(count))
				tmp = append(tmp, string(object[i]))
				break
			}
		} else if i != len(object)-1 && object[i+1] != object[i] {
			tmp = append(tmp, strconv.Itoa(count))
			tmp = append(tmp, string(object[i]))
			count = 1
		}
		if i == len(object)-1 {
			tmp = append(tmp, strconv.Itoa(1))
			tmp = append(tmp, string(object[i]))
		}
	}
	return strings.Join(tmp, "")
}
