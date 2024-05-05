package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	leftSame := isSameTree(p.Left, q.Left)
	rightSame := isSameTree(p.Right, q.Right)
	return leftSame && p.Val == q.Val && rightSame
}
