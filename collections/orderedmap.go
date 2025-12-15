package collections

import (
	"encoding/json"
	"iter"
)

type OrderedMap[K comparable, V any] struct {
	nodes  map[K]*orderedNode[K, V]
	head   *orderedNode[K, V]
	tail   *orderedNode[K, V]
	length int
}

type orderedNode[K comparable, V any] struct {
	key   K
	value V
	prev  *orderedNode[K, V]
	next  *orderedNode[K, V]
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{nodes: make(map[K]*orderedNode[K, V])}
}

func (m *OrderedMap[K, V]) ensure() {
	if m == nil {
		return
	}
	if m.nodes == nil {
		m.nodes = make(map[K]*orderedNode[K, V])
	}
}

// Set inserts or updates the value for the key, preserving insertion order.
func (m *OrderedMap[K, V]) Set(k K, v V) {
	if m == nil {
		return
	}
	m.ensure()
	if node, ok := m.nodes[k]; ok {
		node.value = v
		return
	}
	node := &orderedNode[K, V]{key: k, value: v}
	if m.tail == nil {
		m.head = node
		m.tail = node
	} else {
		m.tail.next = node
		node.prev = m.tail
		m.tail = node
	}
	m.nodes[k] = node
	m.length++
}

func (m *OrderedMap[K, V]) Get(k K) (V, bool) {
	var zero V
	if m == nil || m.nodes == nil {
		return zero, false
	}
	node, ok := m.nodes[k]
	if !ok {
		return zero, false
	}
	return node.value, true
}

func (m *OrderedMap[K, V]) Has(k K) bool {
	if m == nil || m.nodes == nil {
		return false
	}
	_, ok := m.nodes[k]
	return ok
}

// Delete removes the key if present.
func (m *OrderedMap[K, V]) Delete(k K) bool {
	if m == nil || m.nodes == nil {
		return false
	}
	node, ok := m.nodes[k]
	if !ok {
		return false
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		m.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		m.tail = node.prev
	}
	delete(m.nodes, k)
	m.length--
	return true
}

func (m *OrderedMap[K, V]) Keys() []K {
	if m == nil || m.length == 0 {
		return nil
	}
	keys := make([]K, 0, m.length)
	for n := m.head; n != nil; n = n.next {
		keys = append(keys, n.key)
	}
	return keys
}

func (m *OrderedMap[K, V]) Values() []V {
	if m == nil || m.length == 0 {
		return nil
	}
	values := make([]V, 0, m.length)
	for n := m.head; n != nil; n = n.next {
		values = append(values, n.value)
	}
	return values
}

// Range walks the map in insertion order until fn returns false.
func (m *OrderedMap[K, V]) Range(fn func(K, V) bool) {
	if m == nil {
		return
	}
	for n := m.head; n != nil; n = n.next {
		if !fn(n.key, n.value) {
			return
		}
	}
}

// RangeReverse walks the map from newest to oldest until fn returns false.
func (m *OrderedMap[K, V]) RangeReverse(fn func(K, V) bool) {
	if m == nil {
		return
	}
	for n := m.tail; n != nil; n = n.prev {
		if !fn(n.key, n.value) {
			return
		}
	}
}

// All returns an iterator over key-value pairs in insertion order.
func (m *OrderedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m == nil {
			return
		}
		for n := m.head; n != nil; n = n.next {
			if !yield(n.key, n.value) {
				return
			}
		}
	}
}

// Backward returns an iterator over key-value pairs in reverse insertion order.
func (m *OrderedMap[K, V]) Backward() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m == nil {
			return
		}
		for n := m.tail; n != nil; n = n.prev {
			if !yield(n.key, n.value) {
				return
			}
		}
	}
}

func (m *OrderedMap[K, V]) Len() int {
	if m == nil {
		return 0
	}
	return m.length
}

// KeysSlice returns a slice of keys in insertion order.
func (m *OrderedMap[K, V]) KeysSlice() []K {
	return m.Keys()
}

// ValuesSlice returns a slice of values in insertion order.
func (m *OrderedMap[K, V]) ValuesSlice() []V {
	return m.Values()
}

func (m *OrderedMap[K, V]) Clear() {
	if m == nil || m.nodes == nil {
		return
	}
	m.nodes = make(map[K]*orderedNode[K, V])
	m.head = nil
	m.tail = nil
	m.length = 0
}

// MarshalJSON implements the json.Marshaler interface.
// It serializes the OrderedMap as a JSON object, preserving the insertion order
// (though standard JSON parsers may not respect it).
func (m *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	// We build the object manually to control order.
	// Note: This assumes keys are simple strings or can be stringified.
	// For complex keys, standard Go map behaviors apply.
	// For simplicity in this generic implementation, we use a map if keys are strings,
	// but standard library usually iterates maps randomly.
	// To strictly preserve order in output, we must construct the byte slice or use a struct?
	// Actually, Go's json.Marshal sorts map keys.
	// To preserve order, we must output as a JSON object { "k1": v1, "k2": v2 } manually
	// OR output as a list of pairs [ {"Key": k1, "Value": v1}, ... ].
	//
	// Most "Ordered Map" implementations in other langs serialize as an object.
	// Let's implement a manual write to ensure order in the byte stream.

	var buf []byte
	buf = append(buf, '{')
	i := 0
	for k, v := range m.All() {
		if i > 0 {
			buf = append(buf, ',')
		}
		// Marshal Key
		keyBytes, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		// If key was not a string, we might need to conform to JSON object key rules (must be string).
		// json.Marshal will quote strings. If K is int, it outputs number. JSON keys MUST be strings.
		// Detailed generic handling is complex.
		// Fallback: We'll create a temporary map to check how json.Marshal handles K.
		// Better approach for standard lib quality:
		// Just treat it as an object {k:v}. Clients rely on their parser to keep order.
		// We iterate in order, so we emit in order.

		// Note: Simply calling json.Marshal on the map field would lose order.

		// Tricky: K might be int. JSON object keys must be strings.
		// Go's json package handles map[int]T by converting int to string.
		// We replicate that check? Or just delegate to json.Marshal for the key.
		// Let's assume K is string-compatible or user accepts valid JSON key rules.

		// If K is not string, we strip quotes? No, JSON keys must be quoted strings.
		// We will marshal the key, then ensure it is a string.

		if len(keyBytes) > 0 && keyBytes[0] != '"' {
			// Convert non-string key to string for JSON validity
			// e.g. 123 -> "123"
			keyBytes = []byte("\"" + string(keyBytes) + "\"")
		}

		buf = append(buf, keyBytes...)
		buf = append(buf, ':')
		valBytes, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		buf = append(buf, valBytes...)
		i++
	}
	buf = append(buf, '}')
	return buf, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *OrderedMap[K, V]) UnmarshalJSON(data []byte) error {
	// Unmarshaling into an ordered structure is difficult because json.Unmarshal
	// into a map[K]V loses order before we see it.
	// We would need a custom decoder.
	// For now, we fallback to standard map unmarshal, effectively losing order from input.
	// This is a known limitation unless we write a full parser.
	// We will populate the map, order will be random (or specific to Go's map iteration).

	// To support keeping order from JSON, we'd need to parse the token stream.
	// For V1, let's just support loading data.

	tmp := make(map[K]V)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	m.Clear()
	for k, v := range tmp {
		m.Set(k, v)
	}
	return nil
}
