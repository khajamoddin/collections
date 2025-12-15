package concurrent

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkShardedMap_SetGet_StringKeys(b *testing.B) {
	const numShards = 32
	m := NewShardedMap[string, int](numShards, StringHasher)

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			k := "user:" + strconv.Itoa(i)
			m.Set(k, i)
			_, _ = m.Get(k)
			i++
		}
	})
}

func BenchmarkShardedMap_SetGet_Uint64Keys(b *testing.B) {
	const numShards = 32
	m := NewShardedMap[uint64, int](numShards, Uint64Hasher)

	b.RunParallel(func(pb *testing.PB) {
		var i uint64
		for pb.Next() {
			m.Set(i, int(i))
			_, _ = m.Get(i)
			i++
		}
	})
}

func BenchmarkShardedMap_Len(b *testing.B) {
	const (
		numShards = 32
		n         = 100_000
	)

	m := NewShardedMap[int, int](numShards, func(x int) uint64 {
		return Uint64Hasher(uint64(x))
	})

	for i := 0; i < n; i++ {
		m.Set(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Len()
	}
}

func BenchmarkShardedMap_All(b *testing.B) {
	const (
		numShards = 32
		n         = 100_000
	)

	m := NewShardedMap[int, int](numShards, func(x int) uint64 {
		return Uint64Hasher(uint64(x))
	})
	for i := 0; i < n; i++ {
		m.Set(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, _ = range m.All() {
			// just iterate
		}
	}
}

// Baseline map+global-lock comparison
type lockedMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func newLockedMap() *lockedMap {
	return &lockedMap{m: make(map[string]int)}
}

func (m *lockedMap) Set(k string, v int) {
	m.mu.Lock()
	m.m[k] = v
	m.mu.Unlock()
}

func (m *lockedMap) Get(k string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func BenchmarkLockedMap_SetGet_StringKeys(b *testing.B) {
	m := newLockedMap()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			k := "user:" + strconv.Itoa(i)
			m.Set(k, i)
			_, _ = m.Get(k)
			i++
		}
	})
}
