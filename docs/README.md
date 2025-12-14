---
layout: default
title: Overview
nav_order: 2
---

# Collections Go Library

A small, generic collections library for Go providing `Set[T]`, `Deque[T]`, and `PriorityQueue[T]`. Designed to be simple, zero‑value safe where possible, and idiomatic.

- Module: `github.com/khajamoddin/collections`
- Package imports: `github.com/khajamoddin/collections/collections`

## Install

```bash
go get github.com/khajamoddin/collections
```

## Quick Start

```go
package main

import (
    "fmt"
    col "github.com/khajamoddin/collections/collections"
)

func main() {
    // Set
    s := col.NewSet[int]()
    s.Add(1)
    s.Add(2)
    fmt.Println("set len:", s.Len())

    // Deque
    d := col.NewDeque[string]()
    d.PushFront("b")
    d.PushFront("a")
    v, _ := d.PopFront()
    fmt.Println("pop front:", v)

    // PriorityQueue (min-heap)
    pq := col.NewPriorityQueue[int](func(a, b int) bool { return a < b })
    pq.Push(3)
    pq.Push(1)
    pq.Push(2)
    top, _ := pq.Peek()
    fmt.Println("peek:", top)
}
```

## Packages and Types

- `collections.Set[T]`: hash set with add/remove/has/values and set algebra helpers.
- `collections.Deque[T]`: circular-buffer double-ended queue with O(1) amortized operations.
- `collections.PriorityQueue[T]`: heap-backed priority queue with caller-provided comparator.
- `collections.OrderedMap[K,V]`: insertion-ordered map with forward/reverse iteration.
- `collections.MultiMap[K,V]`: one-to-many map with helper methods.

## Examples

- Run: `go run examples/set_example.go`
- The deque example is provided as a function to avoid multiple `main` entries.

## Testing and Benchmarks

- Run tests: `go test ./... -v -cover`
- Benchmarks: `go test -bench=. ./...`

## CI

GitHub Actions workflow at `.github/workflows/go.yml` runs tidy, tests (with coverage), and build on push/PR.

## License

MIT, see `LICENSE`.
