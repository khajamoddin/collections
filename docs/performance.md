---
layout: default
title: Performance Guide
nav_order: 5
---

# Performance Guide

This guide provides complexity guarantees and benchmark results to help you choose the right data structure.

## Big-O Complexity

| Type | Operation | Complexity | Notes |
|------|-----------|------------|-------|
| **Set** | Add | O(1) | Average case (map-backed) |
| | Remove | O(1) | Average case |
| | Has | O(1) | Average case |
| **Deque** | PushFront/Back | O(1) | Amortized (ring buffer) |
| | PopFront/Back | O(1) | Amortized |
| | Random Access | O(1) | Not indexable, but iter is O(N) |
| **OrderedMap** | Set/Get/Delete| O(1) | Average case (map + linked list) |
| | Iteration | O(N) | In insertion order |
| **PriorityQueue**| Push | O(log N) | Standard heap property |
| | Pop | O(log N) | |
| | Peek | O(1) | |

## Benchmarks

Benchmarks were run on Apple M2.

```text
BenchmarkDequePushPop-8         123832432                9.150 ns/op
BenchmarkSetAdd-8               10909752               135.6 ns/op
```

### Interpretation

- **Deque** is extremely fast (9ns/op) because it avoids allocation once the buffer is grown, and mainly involves pointer arithmetic.
- **Set** is comparable to standard Go map insertions (~135ns/op). The wrapper overhead is negligible.
