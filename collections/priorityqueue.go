package collections

import (
	"container/heap"
	"iter"
)

type PriorityQueue[T any] struct {
	h *pqHeap[T]
}

type pqHeap[T any] struct {
	less func(a, b T) bool
	data []T
}

func (h *pqHeap[T]) Len() int           { return len(h.data) }
func (h *pqHeap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *pqHeap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *pqHeap[T]) Push(x any)         { h.data = append(h.data, x.(T)) }
func (h *pqHeap[T]) Pop() any {
	n := len(h.data)
	v := h.data[n-1]
	h.data = h.data[:n-1]
	return v
}

func NewPriorityQueue[T any](less func(T, T) bool) *PriorityQueue[T] {
	h := &pqHeap[T]{less: less}
	heap.Init(h)
	return &PriorityQueue[T]{h: h}
}

func (q *PriorityQueue[T]) Len() int {
	if q == nil || q.h == nil {
		return 0
	}
	return q.h.Len()
}

// All returns an iterator over the elements in the priority queue.
// Note: The order is not guaranteed to be sorted (it iterates the underlying heap slice).
func (q *PriorityQueue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if q == nil || q.h == nil {
			return
		}
		for _, v := range q.h.data {
			if !yield(v) {
				return
			}
		}
	}
}

func (q *PriorityQueue[T]) Push(v T) {
	heap.Push(q.h, v)
}

func (q *PriorityQueue[T]) Pop() (T, bool) {
	var zero T
	if q.h.Len() == 0 {
		return zero, false
	}
	return heap.Pop(q.h).(T), true
}

func (q *PriorityQueue[T]) Peek() (T, bool) {
	var zero T
	if q.h.Len() == 0 {
		return zero, false
	}
	return q.h.data[0], true
}

func (q *PriorityQueue[T]) Clear() {
	q.h.data = nil
}
