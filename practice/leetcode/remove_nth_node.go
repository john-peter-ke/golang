package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. First pass: get the size
	size := 0
	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}

	// Handle removing the head
	if n == size {
		return head.Next
	}

	// 2. Second pass: stop at the node BEFORE the one to remove
	curr = head
	for i := 1; i < size-n; i++ {
		curr = curr.Next
	}

	// 3. Skip the node
	curr.Next = curr.Next.Next

	return head
}

func removeNthFromEndTwoPointer(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
