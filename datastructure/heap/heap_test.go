package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

	t.Run("Test heap", func(t *testing.T) {
		hp := maxHeap{}
		values := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
		for _, v := range values {
			hp.Insert(v)
			fmt.Printf("%v \n", hp)
		}

		for i := 0; i < 5; i++ {
			hp.Extract()
			fmt.Printf("%v \n", hp)
		}
	})
}
