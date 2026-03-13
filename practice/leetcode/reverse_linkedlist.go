package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	current := head
	for current != nil {
		// swap next node to prev
		next := current.Next // place holder for next node
		current.Next = prev  // make next node as previous
		prev = current       // make prev node as current node

		current = next // use this for iteration
	}

	return prev
}
