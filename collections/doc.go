// Package collections provides generic, type-safe data structures for Go.
//
// It fills the gap in the standard library by offering high-quality implementations
// of common structures like Set, Deque, OrderedMap, and PriorityQueue, optimized
// for ergonomics and performance.
//
// # Design Philosophy
//
// - **Zero Dependencies**: Core logic relies only on the Go standard library.
// - **Idiomatic**: APIs allow for fluent, standard Go coding styles.
// - **Performance**: Benchmarked to ensure minimal overhead over raw native maps/slices.
// - **Modern**: Fully supports Go 1.23+ iterators.
//
// # Thread Safety
//
// By default, the collections in this package are **not thread-safe**. This is a deliberate
// choice to avoid the overhead of synchronization for the common single-goroutine use case.
// If you need to access these collections from multiple goroutines, you must guard them
// with a sync.Mutex or RWMutex, or wait for the upcoming concurrent subpackage.
package collections
