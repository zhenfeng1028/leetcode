package medium

import "strings"

func simplifyPath(path string) string {
	if len(path) == 1 {
		return path
	}
	strs := strings.Split(path, "/")
	res := ""
	for _, str := range strs {
		if len(str) == 0 {
			continue
		}
		if str == "." {
			// continue
		} else if str == ".." {
			lastSlashIdx := -1
			for i := len(res) - 1; i > 0; i-- {
				if res[i] == '/' {
					lastSlashIdx = i
					break
				}
			}
			if lastSlashIdx != -1 {
				res = res[:lastSlashIdx]
			} else {
				res = ""
			}
		} else {
			res += "/" + str
		}
	}
	if res == "" {
		return "/"
	}
	return res
}
