package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev

		prev = current
		current = next
	}

	return prev
}
