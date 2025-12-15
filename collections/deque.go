package collections

import "iter"

// Deque is a generic double-ended queue implemented as a ring buffer.
// It supports O(1) amortized push and pop operations at both ends.
type Deque[T any] struct {
	buf  []T
	head int
	size int
}

// NewDeque creates a new empty Deque.
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

// NewDequeWithCapacity creates a new Deque with preallocated capacity.
func NewDequeWithCapacity[T any](capacity int) *Deque[T] {
	if capacity < 0 {
		capacity = 0
	}
	return &Deque[T]{buf: make([]T, capacity)}
}

// Len returns the number of elements in the Deque.
// Complexity: O(1).
func (d *Deque[T]) Len() int {
	if d == nil {
		return 0
	}
	return d.size
}

func (d *Deque[T]) Clear() {
	if d == nil {
		return
	}
	d.buf = nil
	d.head = 0
	d.size = 0
}

func (d *Deque[T]) PushBack(v T) {
	if d == nil {
		return
	}
	d.ensureCapacity(d.size + 1)
	idx := (d.head + d.size) % cap(d.buf)
	d.buf[idx] = v
	d.size++
}

func (d *Deque[T]) PushFront(v T) {
	if d == nil {
		return
	}
	d.ensureCapacity(d.size + 1)
	d.head = (d.head - 1 + cap(d.buf)) % cap(d.buf)
	d.buf[d.head] = v
	d.size++
}

func (d *Deque[T]) PeekFront() (T, bool) {
	var zero T
	if d == nil || d.size == 0 {
		return zero, false
	}
	return d.buf[d.head], true
}

func (d *Deque[T]) PeekBack() (T, bool) {
	var zero T
	if d == nil || d.size == 0 {
		return zero, false
	}
	idx := (d.head + d.size - 1) % cap(d.buf)
	return d.buf[idx], true
}

func (d *Deque[T]) PopFront() (T, bool) {
	var zero T
	if d == nil || d.size == 0 {
		return zero, false
	}
	v := d.buf[d.head]
	d.head = (d.head + 1) % cap(d.buf)
	d.size--
	return v, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	var zero T
	if d == nil || d.size == 0 {
		return zero, false
	}
	idx := (d.head + d.size - 1) % cap(d.buf)
	v := d.buf[idx]
	d.size--
	return v, true
}

// All returns an iterator over elements from front to back.
func (d *Deque[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if d == nil || d.size == 0 {
			return
		}
		for i := 0; i < d.size; i++ {
			idx := (d.head + i) % cap(d.buf)
			if !yield(d.buf[idx]) {
				return
			}
		}
	}
}

// Backward returns an iterator over elements from back to front.
func (d *Deque[T]) Backward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if d == nil || d.size == 0 {
			return
		}
		for i := d.size - 1; i >= 0; i-- {
			idx := (d.head + i) % cap(d.buf)
			if !yield(d.buf[idx]) {
				return
			}
		}
	}
}

func (d *Deque[T]) ensureCapacity(need int) {
	if need <= cap(d.buf) {
		return
	}
	newCap := cap(d.buf)
	if newCap == 0 {
		newCap = 1
	}
	for newCap < need {
		newCap <<= 1
	}
	newBuf := make([]T, newCap)
	if d.size > 0 {
		if d.head+d.size <= cap(d.buf) {
			copy(newBuf, d.buf[d.head:d.head+d.size])
		} else {
			n := cap(d.buf) - d.head
			copy(newBuf, d.buf[d.head:])
			copy(newBuf[n:], d.buf[:d.size-n])
		}
	}
	d.buf = newBuf
	d.head = 0
}
