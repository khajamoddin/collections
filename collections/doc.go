// Package collections provides generic, type-safe data structures for Go.
//
// The standard library offers powerful primitives (maps, slices, channels),
// but many common structures (sets, deques, ordered maps, multimaps, generic
// priority queues) still require repetitive boilerplate or custom code.
//
// collections fills that gap with a small, focused set of types:
//
//   - Set[T]         : generic hash set with set algebra helpers
//   - Deque[T]       : double-ended queue based on a circular buffer
//   - PriorityQueue[T] : generic heap-based priority queue
//   - OrderedMap[K,V]: insertion-ordered map with stable iteration
//   - MultiMap[K,V]  : map from key to multiple values (one-to-many)
//
// # Design goals
//
//   - Idiomatic Go API: methods and naming follow standard library
//     conventions where possible (e.g. Has, Add, Delete, Range).
//   - Zero-value usability: most types are safe to use as their zero value
//     after optional initialization with New… constructors.
//   - Generic and type-safe: all collections are implemented using Go's
//     generics, so you get compile-time type safety without interfaces.
//   - Minimal dependencies: the core package has no external runtime
//     dependencies beyond the Go standard library and, optionally, the
//     iter package when using iterator helpers.
//   - Clarity over cleverness: implementations favor predictable behavior
//     and readable code over micro-optimizations.
//
// # Thread safety
//
// The collections in this package are NOT safe for concurrent use by default.
//
// If a collection is accessed from multiple goroutines, callers must provide
// their own synchronization (for example, using sync.Mutex or sync.RWMutex).
//
// For applications that need built-in synchronization, see the
// "github.com/khajamoddin/collections/collections/concurrent" subpackage,
// which provides concurrency-aware variants such as concurrent sets and
// sharded maps.
//
// # Iterators (Go 1.23+)
//
// When built with Go 1.23 or later, many collections expose iterator
// helpers based on the experimental iter package:
//
//   - Set.All       : func() iter.Seq[T]
//   - OrderedMap.All: func() iter.Seq2[K,V]
//   - MultiMap.All  : func() iter.Seq2[K,V]
//
// These can be used with the Go 1.23 range-over-functions syntax:
//
//	s := collections.NewSet[int]()
//	s.Add(1)
//	s.Add(2)
//
//	for v := range s.All() {
//	    fmt.Println(v)
//	}
//
// On older Go versions, the collections are still usable via their
// standard methods such as Values, Range, Keys, and so on.
//
// # Usage example
//
//	package main
//
//	import (
//	    "fmt"
//
//	    collections "github.com/khajamoddin/collections/collections"
//	)
//
//	func main() {
//	    // Set with set algebra helpers
//	    s1 := collections.NewSet[int]()
//	    s1.Add(1)
//	    s1.Add(2)
//
//	    s2 := collections.NewSet[int]()
//	    s2.Add(2)
//	    s2.Add(3)
//
//	    union := s1.Union(s2) // {1,2,3}
//	    fmt.Println("union:", union.Values())
//
//	    // Ordered map with deterministic iteration
//	    om := collections.NewOrderedMap[string, int]()
//	    om.Set("first", 1)
//	    om.Set("second", 2)
//
//	    om.Range(func(k string, v int) bool {
//	        fmt.Println(k, v)
//	        return true
//	    })
//	}
//
// # Performance characteristics
//
// Each collection documents its complexity guarantees and typical
// performance behavior in its type comment. For example, Set operations
// are O(1) on average, Deque push/pop operations are O(1) amortized,
// and PriorityQueue operations are O(log n).
//
// See the package README and docs/performance.md in the repository
// for benchmark results and more detailed guidance on choosing the
// right collection for a given use case.
package collections
