package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	left, right := 0, len(nums)-1
	mid := left + (right-left)/2
	root := &TreeNode{nums[mid], nil, nil}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
