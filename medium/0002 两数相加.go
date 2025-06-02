package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	prev := ans
	n1, n2 := 0, 0
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		prev.Val = sum % 10
		if l1 != nil || l2 != nil {
			prev.Next = &ListNode{}
			prev = prev.Next
		}
	}
	if carry == 1 {
		prev.Next = &ListNode{1, nil}
	}
	return ans
}
