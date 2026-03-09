package heap

type maxHeap struct {
	array []int
}

func (h *maxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *maxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *maxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}

	maxIndex := 0
	maxValue := h.array[maxIndex]
	size := len(h.array) - 1

	// take last index and put it in the root
	h.array[0] = h.array[size]
	h.array = h.array[:size]

	h.maxHeapifyDown(maxIndex)

	return maxValue
}

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

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
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
