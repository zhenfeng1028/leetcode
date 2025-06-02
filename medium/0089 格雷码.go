package medium

// 计算第 n 个格雷码
// 格雷码计算公式：G(n) = n xor (n >> 1)
func grayCode(n int) []int {
	res := []int{}
	for i := 0; i < 1<<n; i++ {
		res = append(res, i^(i>>1))
	}
	return res
}
