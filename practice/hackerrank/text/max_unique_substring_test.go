package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDistinctSubstringLengthInSessions(t *testing.T) {

	t.Run("MaxDistinctSubstringLengthInSessions", func(t *testing.T) {
		assert.Equal(t, int32(1), maxDistinctSubstringLengthInSessions("aa"))
		// assert.Equal(t, int32(0), maxDistinctSubstringLengthInSessions("*"))
	})
}
