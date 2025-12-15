---
layout: default
title: Roadmap
nav_order: 3
---

# Roadmap

## Current Status (v1.0 Ready)

### Implemented & Stable
- **Core Data Structures**:
  - `Set[T]`: Hash-based set with algebra helpers (Union, Intersection, etc.).
  - `Deque[T]`: Ring-buffer double-ended queue.
  - `PriorityQueue[T]`: Generic wrapper around `container/heap`.
  - `OrderedMap[K,V]`: Insertion-order preserving map.
  - `MultiMap[K,V]`: One-to-many key-value map.
- **Utilities**:
  - `collections/itertools`: Functional helpers for iterators (`Map`, `Filter`, `Reduce`).
  - JSON serialization for `OrderedMap`.
  - Slice conversion helpers (`FromSlice`, `ToSlice`).

### Documentation
- Comprehensive API Reference.
- Performance Guide with Big-O complexity and benchmarks.
- Use-case driven Recipes.
- Integration examples (Redis, Kafka, Kubernetes).

## Planned Features

### Near Term (v1.x)
- **Concurrency**:
  - Thread-safe wrappers (e.g., `SyncSet`, `SyncMap`).
  - Sharded implementations for high-concurrency workloads.
- **Extended Itertools**:
  - `GroupBy`, `Distinct`, `Chunk` helpers.

### Future
- **Analysis Tooling**:
  - CLI to detect suboptimal standard library usage and suggest `collections` replacements.
- **Persistent Data Structures**:
  - Immutable collections for specific functional patterns.
