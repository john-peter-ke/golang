package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {

	t.Run("insert test", func(t *testing.T) {
		word := "aragon"
		trie := NewTrie()
		trie.Insert(word)
		assert.Equal(t, true, trie.Search(word))
	})
}
