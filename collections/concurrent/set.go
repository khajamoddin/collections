package concurrent

import (
	"sync"

	"iter"

	"github.com/khajamoddin/collections/collections"
)

// Set is a concurrency-safe generic set.
//
// It wraps an underlying collections.Set[T] with a RWMutex. All methods
// may be used safely from multiple goroutines.
//
// The zero value is ready to use.
type Set[T comparable] struct {
	mu  sync.RWMutex
	set collections.Set[T]
}

// NewSet constructs an empty concurrent Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{set: *collections.NewSet[T]()}
}

// Add inserts v into the set. It returns true if v was not already present.
// Note: collections.Set.Add returns void, so strictly this wrapper adds locking.
// If the underlying api changes to return bool, we can forward it.
// Actually, the user API requested returning bool. The underlying collections.Set[T].Add() is void.
// We should check Has() first under lock if we generally want bool, but standard collections.Set doesn't return it.
// WAIT: The prompt code "return s.set.Add(v)" assumes Set.Add returns bool.
// Let's check collections.Set.Add signature.
// It is "func (s *Set[T]) Add(v T)". It returns nothing.
// So we must deviate slightly from the prompt or update Set.Add.
// Updating Set.Add is a breaking change? No, adding a return value is mostly compatible but let's stick to the prompt's intent.
// Actually, usually Set.Add returns true if added.
// Let's implement semantics: return true if it wasn't there.
func (s *Set[T]) Add(v T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.set.Has(v) {
		return false
	}
	s.set.Add(v)
	return true
}

// Remove deletes v from the set. It returns true if v was present.
// collections.Set.Remove is void.
func (s *Set[T]) Remove(v T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.set.Has(v) {
		return false
	}
	s.set.Remove(v)
	return true
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Has(v)
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Len()
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Clear()
}

// Values returns a snapshot of the set as a slice.
// The returned slice is not affected by concurrent modifications.
func (s *Set[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]T(nil), s.set.Values()...)
}

// All returns an iterator over a snapshot of the set's values.
//
// The snapshot is taken at the time All is called; concurrent modifications
// after that point are not reflected in the sequence.
func (s *Set[T]) All() iter.Seq[T] {
	values := s.Values()
	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}
