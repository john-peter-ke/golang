package calculation

import (
	"sync"
	"sync/atomic"
)

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func SimpleArraySum(ar []int32) int32 {
	var sum atomic.Int32
	var wg sync.WaitGroup

	for i := 0; i < len(ar); i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			sum.Add(ar[index])
		}(i, &wg)
	}

	wg.Wait()

	return sum.Load()
}
