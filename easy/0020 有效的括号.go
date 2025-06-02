package easy

// 使用栈来匹配
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	hashmap := make(map[string]string)
	hashmap[")"] = "("
	hashmap["]"] = "["
	hashmap["}"] = "{"

	stack := []string{}
	for _, c := range s {
		if v, ok := hashmap[string(c)]; ok {
			if len(stack) <= 0 {
				return false
			}
			if v == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		} else {
			stack = append(stack, string(c))
		}
	}

	return len(stack) == 0
}
