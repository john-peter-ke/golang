package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		word := "aragon"
		trie := NewTrie()
		trie.Insert(word)
		assert.Equal(t, true, trie.Search(word))
	})

	t.Run("Case 2", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("KEYBOARD")
		trie.Insert("KEYCHAIN")
		trie.Insert("KEYNOTE")
		trie.Insert("KETTLE")
		trie.Insert("KITE")

		assert.Equal(t, true, trie.StartWith("KEY"))
		assert.Equal(t, false, trie.StartWith("YEK"))
	})

	t.Run("Case 3", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("KEYBOARD")
		trie.Insert("KEYCHAIN")
		trie.Insert("KEYNOTE")
		trie.Insert("KETTLE")
		trie.Insert("KITE")

		assert.Equal(t, 3, len(trie.Matches("KEY")))
	})
}
