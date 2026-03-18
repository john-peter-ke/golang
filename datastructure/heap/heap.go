package heap

type maxHeap struct {
	array []int
}

// Time Complexity: 0(log n)
// Space Complexity: 0(1)
func (h *maxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *maxHeap) heapifyUp(index int) {
	for h.array[index] > h.array[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

// Time Complexity: 0(log n)
// Space Complexity: 0(1)
func (h *maxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}

	maxIndex := 0
	lastIndex := len(h.array) - 1

	maxValue := h.array[maxIndex]
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.maxHeapifyDown(maxIndex)

	return maxValue
}

// Time Complexity: 0(log n)
// Space Complexity: 0(log n)
func (h *maxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left childer is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		if h.array[childToCompare] > h.array[index] {
			h.swap(childToCompare, index)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}
