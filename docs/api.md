---
layout: default
title: API Reference
nav_order: 4
---

# API Reference

Package: `github.com/khajamoddin/collections/collections`

## Set[T]

- `NewSet[T]() *Set[T]`
- `(*Set[T]) Add(v T)`
- `(*Set[T]) Remove(v T)`
- `(*Set[T]) Has(v T) bool`
- `(*Set[T]) Len() int`
- `(*Set[T]) Clear()`
- `(*Set[T]) Values() []T`

Notes:
- Safe on zero values; internal map is lazily initialized.

## Deque[T]

- `NewDeque[T]() *Deque[T]`
- `(*Deque[T]) PushBack(v T)`
- `(*Deque[T]) PushFront(v T)`
- `(*Deque[T]) PopFront() (T, bool)`
- `(*Deque[T]) PopBack() (T, bool)`
- `(*Deque[T]) PeekFront() (T, bool)`
- `(*Deque[T]) PeekBack() (T, bool)`
- `(*Deque[T]) Len() int`
- `(*Deque[T]) Clear()`

Notes:
- Backed by a slice; operations are amortized efficient for typical usage.

## PriorityQueue[T]

- `NewPriorityQueue[T](less func(T, T) bool) *PriorityQueue[T]`
- `(*PriorityQueue[T]) Push(v T)`
- `(*PriorityQueue[T]) Pop() (T, bool)`
- `(*PriorityQueue[T]) Peek() (T, bool)`
- `(*PriorityQueue[T]) Len() int`
- `(*PriorityQueue[T]) Clear()`

Notes:
- Uses `container/heap` internally.
- Behavior (min/max) depends on the `less` comparator.

## Future Types

- `OrderedMap[K,V]`: maintain insertion order, iterate deterministically.
- `MultiMap[K,V]`: map keys to multiple values efficiently.
