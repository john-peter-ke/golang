package queue

type queue[T any] struct {
	values []T
}

// New creates and returns a new queue
func New[T any]() *queue[T] {
	return &queue[T]{values: make([]T, 0)}
}

// Size returns the length of the queue
func (q *queue[T]) Size() int {
	return len(q.values)
}

// IsEmpty returns true if queue is empty
func (q *queue[T]) IsEmpty() bool {
	return len(q.values) == 0
}

// EnQueue will add value to the end of the queue
func (q *queue[T]) EnQueue(value T) {
	q.values = append(q.values, value)
}

// DeQueue will remove front value added to the queue
func (q *queue[T]) DeQueue() T {
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value
}

// Front returns the first value in the queue
func (q *queue[T]) Front() T {
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue
	}
	return q.values[0]
}

// Rear returns the last value in the queue
func (q *queue[T]) Rear() T {
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue
	}
	return q.values[len(q.values)-1]
}
