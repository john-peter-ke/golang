package tree

import "github.com/mucunga90/go/constraint"

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
