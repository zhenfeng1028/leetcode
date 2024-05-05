package medium

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ans := []string{}
	if n > 1 {
		ss := generateParenthesis(n - 1)
		tmp := "()" + ss[0]
		ans = append(ans, tmp)
		for i := 0; i < len(ss); i++ {
			for j := 0; j < len(ss[i]); j++ {
				if string(ss[i][j]) == "(" {
					tmp = ss[i][:j+1] + "()" + ss[i][j+1:]
					ans = append(ans, tmp)
				}
			}
		}
	}
	return ans
}
