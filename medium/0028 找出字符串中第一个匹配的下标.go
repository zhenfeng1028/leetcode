package medium

func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i <= len(haystack)-n; i++ {
		curIndex := i
		if haystack[i] == needle[0] {
			p := 0
			for i < len(haystack) && p < n {
				if haystack[i] != needle[p] {
					break
				}
				i++
				p++
			}
			if p == n {
				return curIndex
			}
		}
		i = curIndex
	}
	return -1
}
