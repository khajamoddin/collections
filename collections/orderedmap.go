package collections

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

func (m *OrderedMap[K, V]) Len() int {
	if m == nil {
		return 0
	}
	return m.length
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
