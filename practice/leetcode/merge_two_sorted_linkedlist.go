package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tempNode := &ListNode{}
	tail := tempNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	// Attach remaining elements
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}

	return tempNode.Next
}
