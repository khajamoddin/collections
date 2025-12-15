package collections

import "iter"

type MultiMap[K comparable, V comparable] struct {
	m    map[K][]V
	size int
}

func NewMultiMap[K comparable, V comparable]() *MultiMap[K, V] {
	return &MultiMap[K, V]{m: make(map[K][]V)}
}

func (mm *MultiMap[K, V]) ensure() {
	if mm == nil {
		return
	}
	if mm.m == nil {
		mm.m = make(map[K][]V)
	}
}

// Add associates value v with key k.
func (mm *MultiMap[K, V]) Add(k K, v V) {
	if mm == nil {
		return
	}
	mm.ensure()
	mm.m[k] = append(mm.m[k], v)
	mm.size++
}

// Remove removes the first occurrence of v for key k.
func (mm *MultiMap[K, V]) Remove(k K, v V) bool {
	if mm == nil || mm.m == nil {
		return false
	}
	values, ok := mm.m[k]
	if !ok {
		return false
	}
	for i, val := range values {
		if val == v {
			values = append(values[:i], values[i+1:]...)
			mm.size--
			if len(values) == 0 {
				delete(mm.m, k)
			} else {
				mm.m[k] = values
			}
			return true
		}
	}
	return false
}

// RemoveAll removes all values associated with k and returns them.
func (mm *MultiMap[K, V]) RemoveAll(k K) []V {
	if mm == nil || mm.m == nil {
		return nil
	}
	values, ok := mm.m[k]
	if !ok {
		return nil
	}
	delete(mm.m, k)
	mm.size -= len(values)
	return values
}

// Get returns all values for k in insertion order.
func (mm *MultiMap[K, V]) Get(k K) []V {
	if mm == nil || mm.m == nil {
		return nil
	}
	values := mm.m[k]
	if values == nil {
		return nil
	}
	return append([]V(nil), values...)
}

func (mm *MultiMap[K, V]) Has(k K) bool {
	if mm == nil || mm.m == nil {
		return false
	}
	_, ok := mm.m[k]
	return ok
}

// Keys returns the set of keys; order is undefined.
func (mm *MultiMap[K, V]) Keys() []K {
	if mm == nil || len(mm.m) == 0 {
		return nil
	}
	keys := make([]K, 0, len(mm.m))
	for k := range mm.m {
		keys = append(keys, k)
	}
	return keys
}

// Values flattens all stored values in unspecified order.
func (mm *MultiMap[K, V]) Values() []V {
	if mm == nil || len(mm.m) == 0 {
		return nil
	}
	values := make([]V, 0, mm.size)
	for _, vs := range mm.m {
		values = append(values, vs...)
	}
	return values
}

// Len returns the total number of stored values across all keys.
func (mm *MultiMap[K, V]) Len() int {
	if mm == nil {
		return 0
	}
	return mm.size
}

// All returns an iterator over key-value pairs.
func (mm *MultiMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if mm == nil || mm.m == nil {
			return
		}
		for k, values := range mm.m {
			for _, v := range values {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

func (mm *MultiMap[K, V]) Clear() {
	if mm == nil || mm.m == nil {
		return
	}
	mm.m = make(map[K][]V)
	mm.size = 0
}
