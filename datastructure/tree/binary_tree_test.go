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
}
