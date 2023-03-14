package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	num := 0
	prev := head
	for prev != nil {
		num++
		prev = prev.Next
	}
	removeNodeNum := num - n + 1
	if removeNodeNum == 1 {
		return head.Next
	}
	prev = head
	for removeNodeNum-2 > 0 {
		prev = prev.Next
		removeNodeNum--
	}
	prev.Next = prev.Next.Next
	return head
}
