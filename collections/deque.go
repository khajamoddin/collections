package collections

type Deque[T any] struct {
    data []T
}

func NewDeque[T any]() *Deque[T] {
    return &Deque[T]{}
}

func (d *Deque[T]) Len() int {
    return len(d.data)
}

func (d *Deque[T]) Clear() {
    d.data = nil
}

func (d *Deque[T]) PushBack(v T) {
    d.data = append(d.data, v)
}

func (d *Deque[T]) PushFront(v T) {
    d.data = append([]T{v}, d.data...)
}

func (d *Deque[T]) PeekFront() (T, bool) {
    var zero T
    if len(d.data) == 0 {
        return zero, false
    }
    return d.data[0], true
}

func (d *Deque[T]) PeekBack() (T, bool) {
    var zero T
    if len(d.data) == 0 {
        return zero, false
    }
    return d.data[len(d.data)-1], true
}

func (d *Deque[T]) PopFront() (T, bool) {
    var zero T
    if len(d.data) == 0 {
        return zero, false
    }
    v := d.data[0]
    d.data = d.data[1:]
    return v, true
}

func (d *Deque[T]) PopBack() (T, bool) {
    var zero T
    if len(d.data) == 0 {
        return zero, false
    }
    i := len(d.data) - 1
    v := d.data[i]
    d.data = d.data[:i]
    return v, true
}

