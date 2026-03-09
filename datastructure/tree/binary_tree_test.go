package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {

	t.Run("Test binary tree", func(t *testing.T) {
		searchValue := int32(7)
		tree := NewBinaryTree[int32]()

		assert.Equal(t, false, tree.Search(searchValue))
		tree.Insert(100)
		tree.Insert(76)
		tree.Insert(200)
		tree.Insert(searchValue)
		tree.Insert(300)

		assert.Equal(t, 5, tree.Length())
		assert.Equal(t, true, tree.Search(searchValue))
		assert.Equal(t, 3, tree.Depth())
	})

	t.Run("Test binary tree traverse", func(t *testing.T) {
		tree := NewBinaryTree[int32]()
		tree.Insert(100)
		tree.Insert(76)
		tree.Insert(200)
		tree.Insert(7)
		tree.Insert(300)

		assert.Equal(t, []int32{7, 76, 100, 200, 300}, tree.TraverseInorder())
		assert.Equal(t, []int32{100, 76, 7, 200, 300}, tree.TraversePreOrder())
		assert.Equal(t, []int32{7, 76, 300, 200, 100}, tree.TraversePostOrder())
		assert.Equal(t, []int32{100, 76, 200, 7, 300}, tree.TraverseLevelOrder())
	})

	t.Run("Test binary tree invert", func(t *testing.T) {
		tree := NewBinaryTree[int32]()
		tree.Insert(100)
		tree.Insert(76)
		tree.Insert(200)
		tree.Insert(7)
		tree.Insert(300)

		assert.Equal(t, []int32{100, 76, 200, 7, 300}, tree.TraverseLevelOrder())
		tree.InvertTree()
		assert.Equal(t, []int32{100, 200, 76, 300, 7}, tree.TraverseInorder())
	})

	t.Run("Test binary tree insert array", func(t *testing.T) {
		tree := NewBinaryTree[int32]()
		tree.Insert(100)
		tree.Insert(76)
		tree.Insert(200)
		tree.Insert(7)
		tree.Insert(300)

		array := []int32{100, 76, 7, 200, 300}
		assert.Equal(t, tree, NewFromArray(array))
	})
}
