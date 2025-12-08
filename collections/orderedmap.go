package collections

type OrderedMap[K comparable, V any] struct {
	keys []K
	m    map[K]V
}
