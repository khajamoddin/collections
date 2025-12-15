package concurrent

import (
	"sync"

	"iter"
)

// Hasher computes a hash for keys of type K.
//
// Implementations should provide a well-distributed hash for a given K.
// For many applications, simple hashes of strings or integers are sufficient.
type Hasher[K any] func(K) uint64

// shard is a single protected map.
type shard[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

// ShardedMap is a concurrency-safe map[K]V divided into N shards.
//
// Sharding reduces lock contention under concurrent read/write workloads.
// Keys are assigned to shards using the provided Hasher.
type ShardedMap[K comparable, V any] struct {
	shards []shard[K, V]
	hasher Hasher[K]
}

// NewShardedMap creates a new ShardedMap with the given number of shards
// and hash function. numShards must be > 0.
//
// Callers are responsible for choosing an appropriate hash function for K.
// For common key types, see StringHasher and Uint64Hasher.
func NewShardedMap[K comparable, V any](numShards int, hasher Hasher[K]) *ShardedMap[K, V] {
	if numShards <= 0 {
		numShards = 1
	}
	s := &ShardedMap[K, V]{
		shards: make([]shard[K, V], numShards),
		hasher: hasher,
	}
	for i := range s.shards {
		s.shards[i].m = make(map[K]V)
	}
	return s
}

func (m *ShardedMap[K, V]) shardFor(key K) *shard[K, V] {
	h := m.hasher(key)
	idx := int(h % uint64(len(m.shards)))
	return &m.shards[idx]
}

// Set stores the value v for key k.
func (m *ShardedMap[K, V]) Set(k K, v V) {
	s := m.shardFor(k)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

// Get retrieves the value for key k.
//
// The second return value reports whether the key was present.
func (m *ShardedMap[K, V]) Get(k K) (V, bool) {
	s := m.shardFor(k)
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[k]
	return v, ok
}

// Delete removes key k from the map. It reports whether the key was present.
func (m *ShardedMap[K, V]) Delete(k K) bool {
	s := m.shardFor(k)
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.m[k]; !ok {
		return false
	}
	delete(s.m, k)
	return true
}

// Len returns the total number of key/value pairs across all shards.
//
// This call acquires read locks on all shards sequentially.
func (m *ShardedMap[K, V]) Len() int {
	total := 0
	for i := range m.shards {
		s := &m.shards[i]
		s.mu.RLock()
		total += len(s.m)
		s.mu.RUnlock()
	}
	return total
}

// Range calls f for each key/value pair in the map.
//
// Iteration order is not specified. If f returns false, Range stops early.
func (m *ShardedMap[K, V]) Range(f func(K, V) bool) {
	for i := range m.shards {
		s := &m.shards[i]
		s.mu.RLock()
		for k, v := range s.m {
			if !f(k, v) {
				s.mu.RUnlock()
				return
			}
		}
		s.mu.RUnlock()
	}
}

// All returns an iterator over a snapshot of the map's key/value pairs.
//
// The snapshot is taken at the time All is called; concurrent modifications
// after that point are not reflected in the sequence.
func (m *ShardedMap[K, V]) All() iter.Seq2[K, V] {
	// Take a snapshot.
	type kv struct {
		k K
		v V
	}
	var items []kv
	for i := range m.shards {
		s := &m.shards[i]
		s.mu.RLock()
		for k, v := range s.m {
			items = append(items, kv{k: k, v: v})
		}
		s.mu.RUnlock()
	}

	return func(yield func(K, V) bool) {
		for _, it := range items {
			if !yield(it.k, it.v) {
				return
			}
		}
	}
}

// StringHasher is a simple hash function for string keys.
func StringHasher(s string) uint64 {
	// FNV-1a 64-bit
	const (
		offset64 = 1469598103934665603
		prime64  = 1099511628211
	)
	var h uint64 = offset64
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

// Uint64Hasher is a trivial hasher for uint64 keys.
func Uint64Hasher(x uint64) uint64 {
	// Mix bits a bit using a variant of SplitMix64.
	x += 0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	x = x ^ (x >> 31)
	return x
}
