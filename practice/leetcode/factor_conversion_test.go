package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorConvert(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		data1 := []conversion{
			{"mile", "km", 1.60934},
			{"km", "m", 1000},
			{"m", "cm", 100},
		}
		assert.Equal(t, float64(1609.34), factorConvertMaps(data1, "mile", "m"))

		data2 := []conversion{
			{"USD", "EUR", 0.92},
			{"EUR", "GBP", 0.85},
			{"GBP", "JPY", 190.0},
		}
		assert.Equal(t, float64(148.58), factorConvertMaps(data2, "USD", "JPY"))
	})

	t.Run("Case 2", func(t *testing.T) {
		data1 := [][]interface{}{
			{"mile", "km", 1.60934},
			{"km", "m", 1000.0},
			{"m", "cm", 100.0},
		}
		assert.Equal(t, float64(1609.34), factorConvertWithList(data1, "mile", "m"))

		data2 := [][]interface{}{
			{"USD", "EUR", 0.92},
			{"EUR", "GBP", 0.85},
			{"GBP", "JPY", 190.0},
		}
		assert.Equal(t, float64(148.58), factorConvertWithList(data2, "USD", "JPY"))
	})

	t.Run("Case 2", func(t *testing.T) {
		data1 := [][]interface{}{
			{"mile", "km", 1.60934},
			{"km", "m", 1000.0},
			{"m", "cm", 100.0},
		}
		assert.Equal(t, float64(1609.34), factorConvertMatrix(data1, "mile", "m"))

		data2 := [][]interface{}{
			{"USD", "EUR", 0.92},
			{"EUR", "GBP", 0.85},
			{"GBP", "JPY", 190.0},
		}
		assert.Equal(t, float64(148.58), factorConvertMatrix(data2, "USD", "JPY"))
	})
}
