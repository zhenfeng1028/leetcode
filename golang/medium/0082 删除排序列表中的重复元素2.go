package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preHead := &ListNode{0, nil}
	preHead.Next = head
	cur := preHead
	prev := head
	cnt := 1
	for prev.Next != nil {
		if prev.Next.Val == prev.Val {
			cnt++
		} else {
			if cnt == 1 {
				cur.Next = prev
				cur = prev
			} else if cnt > 1 {
				cur.Next = prev.Next
				cnt = 1
			}
		}
		prev = prev.Next
	}
	if cnt > 1 {
		cur.Next = nil
	}
	return preHead.Next
}
