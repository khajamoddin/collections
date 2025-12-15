# collections

[![Go Reference](https://pkg.go.dev/badge/github.com/khajamoddin/collections.svg)](https://pkg.go.dev/github.com/khajamoddin/collections)
[![Go Report Card](https://goreportcard.com/badge/github.com/khajamoddin/collections)](https://goreportcard.com/report/github.com/khajamoddin/collections)
[![CI](https://github.com/khajamoddin/collections/actions/workflows/go.yml/badge.svg)](https://github.com/khajamoddin/collections/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/khajamoddin/collections/branch/main/graph/badge.svg)](https://codecov.io/gh/khajamoddin/collections)

---

## Overview

**collections** is a modern, idiomatic, zero-boilerplate **generic collections library for Go**.  
It provides high-frequency data structures such as:

- `Set[T]`
- `Deque[T]`
- `PriorityQueue[T]`
- `OrderedMap[K,V]`
- `MultiMap[K,V]`

…and utilities for slices, maps, iterators (`collections/itertools`).  

This library is designed for clarity, performance, and Go-style minimalism. It fills a gap in the Go ecosystem by providing **well-tested, reusable generic data structures**.

---

## Features

- **Generic, type-safe collections** leveraging Go 1.18+ generics  
- Zero-value usability wherever possible  
- Clean, minimal API following Go conventions (including set algebra helpers)  
- Ordered iteration via `OrderedMap` and one-to-many with `MultiMap`  
- Fully tested and benchmarked  
- Ready for real-world production use

---

## Installation

### Prerequisites
- Go 1.23 or later (required for iterators)

```bash
go get github.com/khajamoddin/collections
```

---

## Quick Start

```go
package main

import (
	"fmt"

	collections "github.com/khajamoddin/collections/collections"
)

func main() {
	// Set with algebra helpers
	s := collections.NewSet[int]()
	s.Add(1)
	s.Add(2)
	other := collections.NewSet[int]()
	other.Add(2)
	other.Add(3)
	union := s.Union(other) // {1,2,3}
	fmt.Println(union.Values())

	// Iterator support (Go 1.23+)
	for v := range s.All() {
		fmt.Println(v)
	}

	// Deque (circular buffer)
	d := collections.NewDeque[string]()
	d.PushFront("b")
	d.PushBack("c")
	d.PushFront("a")
	v, _ := d.PopFront()
	fmt.Println("front:", v)

	// OrderedMap preserves insertion order
	om := collections.NewOrderedMap[string, int]()
	om.Set("first", 1)
	om.Set("second", 2)
	// Classic Range with closure
	om.Range(func(k string, v int) bool {
		fmt.Println(k, v)
		return true
	})
	// Modern Range loop
	for k, v := range om.All() {
		fmt.Println(k, v)
	}

	// MultiMap stores multiple values per key
	mm := collections.NewMultiMap[string, int]()
	mm.Add("id", 1)
	mm.Add("id", 2)
	fmt.Println("ids:", mm.Get("id"))
	// Iterator
	for k, v := range mm.All() {
		fmt.Println(k, v)
	}

	// PriorityQueue (min-heap)
	pq := collections.NewPriorityQueue[int](func(a, b int) bool { return a < b })
	pq.Push(5)
	pq.Push(1)
pq.Push(3)
peek, _ := pq.Peek()
fmt.Println("top:", peek)
}
```

---

## 🔌 Integration examples

Real-world usage examples live under [`examples/`](./examples):

- `redis-session-index/` – Track active user sessions per region using `Set`.
- `kafka-retry-queue/` – Retry scheduler using `Deque` + `PriorityQueue`.
- `k8s-config-cache/` – In-memory cache for Kubernetes `ConfigMap`s using `OrderedMap` + `MultiMap`.

These are intentionally minimal but realistic patterns you can adapt in your own services.

---

## Why collections?

| Feature            | Standard Library            | This Library                           |
|--------------------|-----------------------------|----------------------------------------|
| Generic Set        | map[T]struct{} boilerplate  | `Set[T]` with set algebra helpers      |
| Deque              | Manual slice gymnastics     | `Deque[T]` circular buffer             |
| Priority Queue     | `container/heap` verbose    | `PriorityQueue[T]` thin wrapper        |
| Ordered Map        | None                        | `OrderedMap[K,V]` preserves order      |
| Multi Map          | map[K][]V (DIY)             | `MultiMap[K,V]` with helpers           |
| Zero-value usable  | Not always                  | Yes, documented                        |

---

## Complexity & Thread Safety (overview)

- `Set[T]`: add/remove/has `O(1)` avg; set algebra clones—`O(n)`; not thread-safe—protect externally for concurrent use.
- `Deque[T]`: push/pop/peek front/back `O(1)` amortized via circular buffer; not thread-safe.
- `PriorityQueue[T]`: push/pop `O(log n)`, peek `O(1)`; not thread-safe.
- `OrderedMap[K,V]`: set/get/delete `O(1)` avg; ordered iteration forward/reverse; not thread-safe.
- `MultiMap[K,V]`: add `O(1)`, remove first match `O(n)` in value slice, get `O(len(values))`; not thread-safe.

### Concurrent collections

The `collections/concurrent` subpackage provides concurrency-safe variants of common patterns.  
They trade a bit of overhead for simpler, correct use from multiple goroutines.

| Type                                   | Thread safety                                      | Typical operations                            | Complexity (per op)                                                         | Notes                                                                 |
|----------------------------------------|----------------------------------------------------|-----------------------------------------------|----------------------------------------------------------------------------|-----------------------------------------------------------------------|
| `concurrent.Set[T]`                    | Safe for concurrent use (internal RWMutex)        | `Add`, `Remove`, `Has`, `Len`, `Values`, `All`| `Add`/`Remove`/`Has` → O(1) avg; `Len` → O(1); `Values`/`All` → O(n) snapshot | Wraps `collections.Set[T]`; `Values`/`All` operate on a snapshot to avoid holding locks during iteration. |
| `concurrent.ShardedMap[K,V]`           | Safe for concurrent use (per-shard RWMutex)       | `Set`, `Get`, `Delete`, `Len`, `Range`, `All` | `Set`/`Get`/`Delete` → O(1) avg; `Len` → O(S + n) where S = shard count; `Range`/`All` → O(n) | Hash-based sharding reduces contention under mixed read/write workloads; iteration order is undefined. |

**Notes**

- All concurrent types are designed for **many readers and writers** in typical backend workloads.
- For read-mostly maps, using more shards (e.g. 32–128) can reduce lock contention further.
- For latency-sensitive paths, you should still **profile** and tune shard counts or use workload-specific designs.
