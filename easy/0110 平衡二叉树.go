package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(height(root.Left)-height(root.Right)) <= 1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
