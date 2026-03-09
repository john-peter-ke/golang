package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHastTable(t *testing.T) {

	t.Run("test hashtable", func(t *testing.T) {
		var key1 string = "ERIC"
		var key2 string = "KENNY"
		var key3 string = "KYLE"
		var key4 string = "STAN"
		var key5 string = "RANDY"
		var key6 string = "BUTTERS"
		var key7 string = "TOKEN"
		hashtable := NewHashTable()
		hashtable.Insert(key1)
		hashtable.Insert(key2)
		hashtable.Insert(key3)
		hashtable.Insert(key4)
		hashtable.Insert(key5)
		hashtable.Insert(key6)
		hashtable.Insert(key7)
		assert.Equal(t, true, hashtable.Search(key2))
		assert.Equal(t, true, hashtable.Delete(key2))
		assert.Equal(t, false, hashtable.Search(key2))
	})
}
