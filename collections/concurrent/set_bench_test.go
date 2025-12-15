package concurrent

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkConcurrentSet_Add(b *testing.B) {
	s := NewSet[int]()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			s.Add(i)
			i++
		}
	})
}

func BenchmarkConcurrentSet_Has(b *testing.B) {
	const n = 100_000

	s := NewSet[int]()
	for i := 0; i < n; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = s.Has(i % n)
			i++
		}
	})
}

func BenchmarkConcurrentSet_Values(b *testing.B) {
	const n = 100_000

	s := NewSet[int]()
	for i := 0; i < n; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Values()
	}
}

// Optional: compare with a manual map+mutex implementation
type mapSet struct {
	mu sync.RWMutex
	m  map[int]struct{}
}

func newMapSet() *mapSet {
	return &mapSet{m: make(map[int]struct{})}
}

func (s *mapSet) Add(v int) {
	s.mu.Lock()
	s.m[v] = struct{}{}
	s.mu.Unlock()
}

func (s *mapSet) Has(v int) bool {
	s.mu.RLock()
	_, ok := s.m[v]
	s.mu.RUnlock()
	return ok
}

func BenchmarkMapSet_Add(b *testing.B) {
	s := newMapSet()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			s.Add(i)
			i++
		}
	})
}

func BenchmarkMapSet_Has(b *testing.B) {
	const n = 100_000

	s := newMapSet()
	for i := 0; i < n; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = s.Has(i % n)
			i++
		}
	})
}

// String-based variant for workloads with string keys.
func BenchmarkConcurrentSet_Add_String(b *testing.B) {
	s := NewSet[string]()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			s.Add(strconv.Itoa(i))
			i++
		}
	})
}
