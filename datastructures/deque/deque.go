package deque

type deque[T any] struct {
	values []T
}

// New creates and returns a new deque
func NewDoublyEndedQueue[T any]() *deque[T] {
	return &deque[T]{values: make([]T, 0)}
}

// Size returns the length of the deque
func (d *deque[T]) Size() int {
	return len(d.values)
}

// IsEmpty returns true if deque is empty
func (d *deque[T]) IsEmpty() bool {
	return len(d.values) == 0
}

// EnQueueFront will add item to the front of the deque
func (d *deque[T]) EnQueueFront(item T) {
	d.values = append([]T{item}, d.values...)
}

// EnQueueRear will add item to the back of the deque
func (d *deque[T]) EnQueueRear(item T) {
	d.values = append(d.values, item)
}

// DeQueueFront will remove first item added to the deque
func (d *deque[T]) DeQueueFront() T {
	var zeroValue T
	if d.IsEmpty() {
		return zeroValue
	}
	item := d.values[0]
	d.values = d.values[1:]
	return item
}

// DeQueueRear will remove last item added to the deque
func (d *deque[T]) DeQueueRear() T {
	var zeroValue T
	if d.IsEmpty() {
		return zeroValue
	}
	index := d.Size() - 1
	item := d.values[index]
	d.values = d.values[:index]
	return item
}

// Front returns the first item in the queue
func (d *deque[T]) Front() T {
	var zeroValue T
	if d.IsEmpty() {
		return zeroValue
	}
	return d.values[0]
}

// Rear returns the last item in the queue
func (d *deque[T]) Rear() T {
	var zeroValue T
	if d.IsEmpty() {
		return zeroValue
	}
	return d.values[len(d.values)-1]
}
