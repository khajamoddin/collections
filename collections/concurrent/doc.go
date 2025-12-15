// Package concurrent provides concurrency-safe variants of common collections.
//
// The types in the root collections package are not safe for concurrent use
// without external synchronization. This subpackage wraps those patterns with
// internal locking and, for maps, sharding strategies to reduce contention.
//
// The goal is not to replace careful design of shared state, but to provide
// simple, safe building blocks for common patterns such as:
//
//   - Shared sets used by multiple worker goroutines
//   - Concurrent caches and registries
//   - Sharded maps for high read/write throughput
//
// Types in this package favor predictable behavior and clarity over
// lock-free or highly specialized algorithms. Where performance is
// critical, callers should still profile and, if necessary, tailor
// data structures to their specific workload.
package concurrent
