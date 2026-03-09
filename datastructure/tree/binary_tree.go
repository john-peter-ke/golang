package tree

import (
	"github.com/mucunga90/go/constraint"
)

type btree[T constraint.Ordered] struct {
	root *bnode[T]
}

type bnode[T constraint.Ordered] struct {
	value       T
	left, right *bnode[T]
}

func NewBinaryTree[T constraint.Ordered]() *btree[T] {
	return &btree[T]{}
}

func (n *btree[T]) Insert(v T) {
	if n.root == nil {
		n.root = &bnode[T]{value: v}
		return
	}

	n.root.insert(v)
}

func (n *bnode[T]) insert(v T) {
	if v > n.value {
		// move to right
		if n.right == nil {
			n.right = &bnode[T]{
				value: v,
			}
		} else {
			n.right.insert(v)
		}
	} else if v < n.value {
		// move to left
		if n.left == nil {
			n.left = &bnode[T]{
				value: v,
			}
		} else {
			n.left.insert(v)
		}
	}
}

func (n *btree[T]) Search(v T) bool {
	if n.root == nil {
		return false
	}

	return n.root.search(v)
}

func (n *bnode[T]) search(v T) bool {
	if n == nil {
		return false
	}

	if v > n.value {
		// move right
		return n.right.search(v)
	} else if v < n.value {
		// move left
		return n.left.search(v)
	}

	return true
}

func (n *btree[T]) Length() int {
	count := 0
	if n.root == nil {
		return count
	}

	var countNode func(n *bnode[T])
	countNode = func(n *bnode[T]) {
		if n != nil {
			count++
			countNode(n.left)
			countNode(n.right)
		}
	}

	countNode(n.root)
	return count
}

func (n *btree[T]) Depth() int {
	var maxLength func(n *bnode[T]) int
	maxLength = func(n *bnode[T]) int {
		if n == nil {
			return 0
		}

		rightLen := maxLength(n.right)
		leftLen := maxLength(n.left)

		if rightLen > leftLen {
			return rightLen + 1
		}

		return leftLen + 1
	}

	return maxLength(n.root)
}

// DFS Traversals
// Inorder (Left → Root → Right)
// Used in Binary Search Trees to get sorted order.
func (n *btree[T]) TraverseInorder() []T {
	nodes := make([]T, 0)
	var traverse func(node *bnode[T])
	traverse = func(node *bnode[T]) {
		if node == nil {
			return
		}

		traverse(node.left)
		nodes = append(nodes, node.value)
		traverse(node.right)
	}

	traverse(n.root)
	return nodes
}

// DFS Traversals
// Preorder (Root → Left → Right)
// Used for tree copying and serialization.
func (n btree[T]) TraversePreOrder() []T {
	nodes := make([]T, 0)
	var traverse func(node *bnode[T])
	traverse = func(node *bnode[T]) {
		if node == nil {
			return
		}

		nodes = append(nodes, node.value)
		traverse(node.left)
		traverse(node.right)
	}

	traverse(n.root)
	return nodes
}

// DFS Traversals
// Postorder (Left → Right → Root)
// Used when deleting trees.
func (n btree[T]) TraversePostOrder() []T {
	nodes := make([]T, 0)
	var traverse func(node *bnode[T])
	traverse = func(node *bnode[T]) {
		if node == nil {
			return
		}

		traverse(node.left)
		traverse(node.right)
		nodes = append(nodes, node.value)
	}

	traverse(n.root)
	return nodes
}

// BFS Traversals
func (n btree[T]) TraverseLevelOrder() []T {
	nodes := make([]T, 0)
	if n.root == nil {
		return nodes
	}

	queue := make([]*bnode[T], 0)
	queue = append(queue, n.root)
	for len(queue) > 0 {
		node := queue[0]
		nodes = append(nodes, node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}

		queue = queue[1:]
	}

	return nodes
}

func (n btree[T]) InvertTree() {
	var invert func(node *bnode[T])
	invert = func(node *bnode[T]) {
		if node == nil {
			return
		}

		ln := node.left
		rn := node.right
		node.left = rn
		node.right = ln
		invert(node.left)
		invert(node.right)
	}

	invert(n.root)
}

func NewFromArray[T constraint.Ordered](elements []T) *btree[T] {
	tree := &btree[T]{}
	tree.root = addFromSortedArray(elements, 0, len(elements)-1)
	return tree
}

func addFromSortedArray[T constraint.Ordered](elements []T, start, end int) *bnode[T] {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &bnode[T]{value: elements[mid]}
	node.left = addFromSortedArray(elements, start, mid-1)
	node.right = addFromSortedArray(elements, mid+1, end)

	return node
}
