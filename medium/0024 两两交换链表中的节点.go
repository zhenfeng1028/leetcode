package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, cur := head, head.Next
	// 用于返回
	firstNode := &ListNode{}
	// 用于交换后与前面的节点连接
	connectNode := &ListNode{}
	firstTwoNodes := true
	for prev != nil && cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev.Next = tmp
		connectNode.Next = cur
		connectNode = prev
		if firstTwoNodes {
			firstNode = cur
			firstTwoNodes = false
		}
		prev = prev.Next
		if tmp != nil {
			cur = tmp.Next
		}
	}
	return firstNode
}
