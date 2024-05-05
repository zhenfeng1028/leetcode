package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	leftRes := inorderTraversal(root.Left)
	rightRes := inorderTraversal(root.Right)
	ans := []int{}
	ans = append(ans, leftRes...)
	ans = append(ans, root.Val)
	ans = append(ans, rightRes...)
	return ans
}
