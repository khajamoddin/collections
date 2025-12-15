# Performance

The goal of `collections` is not to win every micro-benchmark, but to provide
predictable, well-behaved data structures that are "fast enough" for real
production systems while dramatically reducing boilerplate and bugs.

This document summarizes the performance characteristics we care about and
how the library behaves in practice.

---

## Core collections

The core types (`Set`, `Deque`, `PriorityQueue`, `OrderedMap`, `MultiMap`)
are thin, generic wrappers around patterns that Go developers already use:

- `Set[T]` → `map[T]struct{}` + set algebra helpers
- `Deque[T]` → ring-buffer over a slice
- `PriorityQueue[T]` → heap-based queue
- `OrderedMap[K,V]` → linked list + `map[K]*node`
- `MultiMap[K,V]` → `map[K][]V` + helpers

As a result, the complexity guarantees match what you would expect:

- `Set.Add`/`Has`/`Remove` → O(1) average
- `Deque.PushFront`/`PushBack`/`PopFront`/`PopBack` → O(1) amortized
- `PriorityQueue.Push`/`Pop` → O(log n), `Peek` → O(1)
- `OrderedMap.Set`/`Get`/`Delete` → O(1) average; iteration → O(n)
- `MultiMap.Add` → O(1); `Get` → O(len(values)) for that key

In practice, benchmarks show that using `collections` instead of hand-written
`slices` and `maps` introduces negligible overhead while giving you:

- Safer semantics (e.g. real set algebra instead of ad-hoc loops)
- Cleaner, more expressive code
- Less duplication across services

If you are currently maintaining your own `Set` / `Deque` / heap wrappers in
multiple services, `collections` should be a drop-in improvement.

---

## Concurrent collections

The `collections/concurrent` subpackage adds correctness under concurrency
by wrapping common patterns in locks and shards:

- `concurrent.Set[T]` wraps `collections.Set[T]` with a `sync.RWMutex`.
- `concurrent.ShardedMap[K,V]` splits a `map[K]V` across N shards, each
  protected by its own `sync.RWMutex`.

The complexity remains the same in big-O terms, but lock contention changes
the real-world performance story.

### Why sharding helps

A single `map+RWMutex` works well at low concurrency: every operation grabs
the same lock, does O(1) work, and releases it. Under many goroutines doing
mixed read/write workloads, that single lock becomes a hotspot:

- Writers block readers.
- Readers briefly block other readers.
- Latency tails get worse as more goroutines pile up.

`ShardedMap` spreads keys across multiple independent maps and locks:

- Different keys often hit different shards → operations proceed in parallel.
- Contention on a single lock only affects a fraction of the keys.
- Tail latencies are typically lower under realistic load.

Benchmarks (Apple M2, 8 cores, 32 shards for ShardedMap):

| Benchmark                                | ns/op | B/op | allocs/op |
|------------------------------------------|-------|------|-----------|
| `BenchmarkConcurrentSet_Add`             | 291   | 8    | 0         |
| `BenchmarkMapSet_Add` (Baseline)         | 365   | 7    | 0         |
| `BenchmarkShardedMap_SetGet_StringKeys`  | 262   | 33   | 2         |
| `BenchmarkLockedMap_SetGet_StringKeys`   | 659   | 33   | 2         |

**Analysis**:
- `ShardedMap` is **~2.5x faster** than a global lock map (`LockedMap`) under high contention (`SetGet` mixed workload).
- `ConcurrentSet` adds minimal overhead vs a raw mutex-protected map, while offering a richer API.

### When to use concurrent types

1. **Use core collections** when:
   - You manipulate them from a single goroutine (or guard them yourself).
   - You are writing tight, hot-path code and want full manual control.

2. **Use concurrent.Set / ShardedMap** when:
   - Multiple goroutines need to share and update a set or map.
   - You want correctness and simplicity without designing your own sharding or lock strategy.
   - You are willing to trade a small constant-factor overhead for safer, easier-to-maintain code.

---

## Iterators and itertools

Go 1.23 introduces range-over-functions and the `iter` package, which makes it
possible to express higher-level data pipelines in a style similar to LINQ or
Java streams.

The `collections/itertools` subpackage provides helpers such as:
- `Map` – transform a sequence
- `Filter` – keep only values that match a predicate
- `Reduce` – fold a sequence into a single value
- `ToSlice` – materialize a sequence into a slice

### Iterator overhead in practice

The natural concern is: "aren't these abstractions slower than a plain for?"

Benchmarks (100k elements):

| Benchmark (Iter vs Slice)         | Iter (ns/op) | Slice (ns/op) | Diff    |
|-----------------------------------|--------------|---------------|---------|
| `Map_IntToInt`                    | 108,999      | 102,567       | +6%     |
| `Filter_Even`                     | 112,492      | 71,143        | +58%    |
| `Reduce_Sum`                      | 57,470       | 63,437        | -10%    |
| `MapFilter_Pipeline`              | 161,525      | 129,530       | +24%    |

**Analysis**:

- Iterator-based pipelines do introduce a small overhead per element, mainly due to function calls and indirection (esp. in `Filter`).
- However, `Reduce` via iterators was actually *faster* in this run, likely due to better inlining or bounds-check elimination patterns in the iterator producer.
- For most backend workloads (I/O bounded), this 6-25% CPU overhead on simple loops is negligible compared to network calls.

**Recommendation**:
- Use iterators for readability in business logic, data transformation pipelines, and complex filtering.
- Stick to manual loops in hot-path numeric kernels or where every nanosecond counts.

---

## How to choose the right building block

As a rule of thumb:

1. Start with the simplest **core collection** (`Set`, `Deque`, `PriorityQueue`, `OrderedMap`) that fits your mental model.
2. Reach for **concurrent variants** when correctness under concurrency matters more than raw single-thread speed.
3. Use **iterators** when you value composability and readability over micro-optimized lookups.

Measure on your own hardware, in your own workloads, and treat benchmarks as guidance rather than gospel. `collections` aims to give you sharp, ergonomic tools so you can focus on system design.
