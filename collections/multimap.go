package collections

type MultiMap[K comparable, V any] struct {
	m map[K][]V
}
