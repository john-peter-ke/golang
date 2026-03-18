package leetcode

func linkedListHasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	nodes := make(map[*ListNode]*ListNode)
	curr := head
	for curr != nil {
		if _, ok := nodes[curr]; ok {
			return true
		}

		nodes[curr] = curr
		curr = curr.Next
	}

	return false
}

func linkedListHasCycleTwoPointer(head *ListNode) bool {
	// 1. Edge case: empty list or single node cannot have a cycle
	if head == nil || head.Next == nil {
		return false
	}

	// 2. Initialize two pointers
	slow := head      // moves 1 step
	fast := head.Next // moves 2 steps
	// 3. Traverse the list
	for slow != fast {
		// If fast hits the end (nil), there is no cycle
		if fast == nil || fast.Next == nil {
			return false
		}

		// Move pointers
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 4. If the loop breaks, it means slow == fast (they met!)
	return true
}
