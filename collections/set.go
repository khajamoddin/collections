package collections

type Set[T comparable] struct {
    m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) ensure() {
    if s.m == nil {
        s.m = make(map[T]struct{})
    }
}

func (s *Set[T]) Add(v T) {
    s.ensure()
    s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
    if s.m == nil {
        return
    }
    delete(s.m, v)
}

func (s *Set[T]) Has(v T) bool {
    if s.m == nil {
        return false
    }
    _, ok := s.m[v]
    return ok
}

func (s *Set[T]) Len() int {
    return len(s.m)
}

func (s *Set[T]) Clear() {
    if s.m == nil {
        return
    }
    s.m = make(map[T]struct{})
}

func (s *Set[T]) Values() []T {
    if s.m == nil {
        return nil
    }
    out := make([]T, 0, len(s.m))
    for k := range s.m {
        out = append(out, k)
    }
    return out
}

