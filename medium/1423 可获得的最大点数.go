package medium

/*
	几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。

	每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。

	你的点数就是你拿到手中的所有卡牌的点数之和。

	给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。
*/

// 方法一：数组总和 - (长度为 n-k 的滑动窗口最小值)
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	// 滑动窗口大小为 n-k
	windowSize := n - k
	// 选前 n-k 个作为初始值
	sum := 0
	for _, pt := range cardPoints[:windowSize] {
		sum += pt
	}
	minSum := sum
	for i := windowSize; i < n; i++ {
		// 滑动窗口每向右移动一格，增加从右侧进入窗口的元素值，并减少从左侧离开窗口的元素值
		sum += cardPoints[i] - cardPoints[i-windowSize]
		minSum = min(minSum, sum)
	}
	total := 0
	for _, pt := range cardPoints {
		total += pt
	}
	return total - minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二：直接滑
func maxScore2(cardPoints []int, k int) int {
	n := len(cardPoints)
	// 选前 k 个作为初始值
	sum := 0
	for _, pt := range cardPoints[:k] {
		sum += pt
	}
	maxSum := sum
	for i := 1; i <= k; i++ {
		// 滑动窗口每向左移动一格，增加从左侧进入窗口的元素值，并减少从右侧离开窗口的元素值
		sum += cardPoints[n-i] - cardPoints[k-i]
		maxSum = max(maxSum, sum)
	}

	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
