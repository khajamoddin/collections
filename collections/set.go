package collections

import "iter"

type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet creates a new empty Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// NewSetFromSlice creates a new Set containing the elements of the slice.
func NewSetFromSlice[T comparable](s []T) *Set[T] {
	set := NewSetWithCapacity[T](len(s))
	for _, v := range s {
		set.Add(v)
	}
	return set
}

// ToSlice returns a slice containing the elements of the Set.
// The order of elements is undefined.
func (s *Set[T]) ToSlice() []T {
	return s.Values()
}

// NewSetWithCapacity creates a set with space preallocated for the given
// capacity. It remains safe to use a zero-value Set without calling this.
func NewSetWithCapacity[T comparable](capacity int) *Set[T] {
	if capacity < 0 {
		capacity = 0
	}
	return &Set[T]{m: make(map[T]struct{}, capacity)}
}

func (s *Set[T]) ensure() {
	if s == nil {
		return
	}
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
}

func (s *Set[T]) Add(v T) {
	if s == nil {
		return
	}
	s.ensure()
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	if s == nil || s.m == nil {
		return
	}
	delete(s.m, v)
}

func (s *Set[T]) Has(v T) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int {
	if s == nil || s.m == nil {
		return 0
	}
	return len(s.m)
}

func (s *Set[T]) Clear() {
	if s == nil || s.m == nil {
		return
	}
	s.m = make(map[T]struct{})
}

// All returns an iterator over the elements in the set.
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if s == nil || s.m == nil {
			return
		}
		for k := range s.m {
			if !yield(k) {
				return
			}
		}
	}
}

func (s *Set[T]) Values() []T {
	if s == nil || s.m == nil {
		return nil
	}
	out := make([]T, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

// Clone returns a shallow copy of the set.
func (s *Set[T]) Clone() *Set[T] {
	if s == nil || s.m == nil {
		return NewSet[T]()
	}
	out := NewSetWithCapacity[T](s.Len())
	for v := range s.m {
		out.m[v] = struct{}{}
	}
	return out
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	if s == nil || s.m == nil {
		return other.Clone()
	}
	otherLen := 0
	if other != nil {
		otherLen = other.Len()
	}
	out := NewSetWithCapacity[T](s.Len() + otherLen)
	for v := range s.m {
		out.m[v] = struct{}{}
	}
	if other != nil {
		for v := range other.m {
			out.m[v] = struct{}{}
		}
	}
	return out
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	if s == nil || other == nil {
		return out
	}
	if s.Len() > other.Len() {
		s, other = other, s
	}
	for v := range s.m {
		if _, ok := other.m[v]; ok {
			out.m[v] = struct{}{}
		}
	}
	return out
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	if s == nil {
		return out
	}
	for v := range s.m {
		if other == nil {
			out.m[v] = struct{}{}
			continue
		}
		if _, ok := other.m[v]; !ok {
			out.m[v] = struct{}{}
		}
	}
	return out
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	if s != nil {
		for v := range s.m {
			if other == nil {
				out.m[v] = struct{}{}
				continue
			}
			if _, ok := other.m[v]; !ok {
				out.m[v] = struct{}{}
			}
		}
	}
	if other != nil {
		for v := range other.m {
			if s == nil {
				out.m[v] = struct{}{}
				continue
			}
			if _, ok := s.m[v]; !ok {
				out.m[v] = struct{}{}
			}
		}
	}
	return out
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s == nil {
		return true
	}
	if other == nil && s.Len() > 0 {
		return false
	}
	for v := range s.m {
		if _, ok := other.m[v]; !ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	if s == nil || other == nil {
		return true
	}
	if s.Len() > other.Len() {
		s, other = other, s
	}
	for v := range s.m {
		if _, ok := other.m[v]; ok {
			return false
		}
	}
	return true
}
