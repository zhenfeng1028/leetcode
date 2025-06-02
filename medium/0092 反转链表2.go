package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	count := 1
	prev := head
	// tmp1为left节点，tmp2为right+1节点
	tmp1, tmp2 := &ListNode{}, &ListNode{}
	for prev != nil {
		if count == left {
			tmp1 = prev
		}
		if count == right+1 {
			tmp2 = prev
			break
		}
		prev = prev.Next
		count++
	}
	// 表明right为最后一个节点
	if prev == nil {
		tmp2 = nil
	}
	prev2 := tmp1
	count = right - left
	for count > 0 {
		prev2 = prev2.Next
		count--
	}
	prev2.Next = nil
	// 调用链表反转方法
	tmp := reverseList(tmp1)
	if left != 1 {
		prev3 := head
		for left-2 > 0 {
			prev3 = prev3.Next
			left--
		}
		prev3.Next = tmp
	} else {
		head = tmp
	}
	prev4 := head
	for prev4 != nil {
		if prev4.Next == nil {
			prev4.Next = tmp2
			break
		}
		prev4 = prev4.Next
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	cur := head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	head.Next = nil
	return prev
}
