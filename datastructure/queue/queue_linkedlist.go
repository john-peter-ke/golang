package queue

// node will store value and next node as well
type node[T any] struct {
	Value T
	Next  *node[T]
}

// queueNode structure with head, tail and length
type queueNode[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// NewQueueNode creates and returns queue
func NewQueueNode[T any]() *queueNode[T] {
	return &queueNode[T]{head: nil, tail: nil, length: 0}
}

// IsEmpty returns true if queue is empty
func (q *queueNode[T]) IsEmpty() bool {
	return q.length == 0
}

// Size returns number of items in queue
func (q *queueNode[T]) Size() int {
	return q.length
}

// 