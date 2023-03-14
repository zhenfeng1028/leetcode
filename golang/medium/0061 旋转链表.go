package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	prev := head
	n := 1
	for prev.Next != nil {
		n++
		prev = prev.Next
	}
	add := n - k%n
	if add == n {
		return head
	}
	prev.Next = head // 首尾相接
	for add > 0 {
		prev = prev.Next
		add--
	}
	ret := prev.Next
	prev.Next = nil
	return ret
}
