package tree

import "github.com/mucunga90/go/constraint"

type tnode[T constraint.Ordered] struct {
	value    T
	children []*tnode[T]
}

func NewTree[T constraint.Ordered](value T) *tnode[T] {
	return &tnode[T]{value: value}
}

// func (t *tnode[T]) Insert(v T) {
// 	if t.value == nil {
// 		t.value = v
// 		return
// 	}
// 	t.insert(v, t.children)
// }

// func (t *tnode[T]) insert(v T, children []*tnode[T]) {
// 	for k, val := range children {
// 		if v > val.value
// 	}
// }
