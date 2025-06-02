package medium

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	hashmap := make(map[string][]string)
	hashmap["2"] = []string{"a", "b", "c"}
	hashmap["3"] = []string{"d", "e", "f"}
	hashmap["4"] = []string{"g", "h", "i"}
	hashmap["5"] = []string{"j", "k", "l"}
	hashmap["6"] = []string{"m", "n", "o"}
	hashmap["7"] = []string{"p", "q", "r", "s"}
	hashmap["8"] = []string{"t", "u", "v"}
	hashmap["9"] = []string{"w", "x", "y", "z"}
	if len(digits) == 1 {
		if letters, ok := hashmap[string(digits[0])]; ok {
			return letters
		}
	}
	tmp := [][]string{}
	for _, v := range digits {
		if letters, ok := hashmap[string(v)]; ok {
			tmp = append(tmp, letters)
		}
	}
	ans := []string{}
	for _, v := range tmp[0] {
		for _, vv := range tmp[1] {
			if len(tmp) > 2 {
				for _, vvv := range tmp[2] {
					if len(tmp) > 3 {
						for _, vvvv := range tmp[3] {
							ans = append(ans, v+vv+vvv+vvvv)
						}
					} else {
						ans = append(ans, v+vv+vvv)
					}
				}
			} else {
				ans = append(ans, v+vv)
			}
		}
	}
	return ans
}
