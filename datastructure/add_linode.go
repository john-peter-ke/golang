package leetcode

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	return createSumNode(traverseNode(l1), traverseNode(l2))
// }

// func createSumNode(n1, n2 []int) *ListNode {
// 	var sum int
// 	maxLen := max(len(n1), len(n2))
// 	for i := maxLen - 1; i > 0; i-- {
// 		var val1, val2 int
// 		if len(n1) > i {
// 			val1 = n1[i]
// 		}

// 		if len(n2) > i {
// 			val1 = n2[i]
// 		}

// 		sum += (val1 + val2)
// 	}

// 	return &ListNode{
// 		Val: sum,
// 	}
// }

// func traverseNode(node *ListNode) []int {
// 	var num []int
// 	if node != nil {
// 		for node.Next != nil {
// 			num = append(num, node.Val)
// 			node = node.Next
// 		}
// 	}
// 	return num
// }
