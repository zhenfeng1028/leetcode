package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	node := &ListNode{}
	if list1.Val < list2.Val {
		node.Val = list1.Val
		node.Next = mergeTwoLists(list1.Next, list2)
	} else {
		node.Val = list2.Val
		node.Next = mergeTwoLists(list1, list2.Next)
	}
	return node
}

// 迭代
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 哨兵节点
	prehead := &ListNode{}
	// 维护一个移动指针
	prev := prehead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 == nil {
		prev.Next = list2
	}
	if list2 == nil {
		prev.Next = list1
	}
	return prehead.Next
}
