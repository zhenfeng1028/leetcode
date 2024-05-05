package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	small := &ListNode{0, nil}
	large := &ListNode{0, nil}
	prev1, prev2 := small, large
	for head != nil {
		if head.Val < x {
			prev1.Next = head
			prev1 = prev1.Next
		} else {
			prev2.Next = head
			prev2 = prev2.Next
		}
		head = head.Next
	}
	if small.Next != nil {
		if large.Next != nil {
			prev1.Next = large.Next
			prev2.Next = nil
		} else {
			prev1.Next = nil
		}
		return small.Next
	}
	prev2.Next = nil
	return large.Next
}
