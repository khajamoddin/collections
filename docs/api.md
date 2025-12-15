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
- `(*Set[T]) All() iter.Seq[T]`
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
- `(*Deque[T]) All() iter.Seq[T]`
- `(*Deque[T]) Backward() iter.Seq[T]`
- `(*Deque[T]) Clear()`

Notes:
- Backed by a slice; operations are amortized efficient for typical usage.

## PriorityQueue[T]

- `NewPriorityQueue[T](less func(T, T) bool) *PriorityQueue[T]`
- `(*PriorityQueue[T]) Push(v T)`
- `(*PriorityQueue[T]) Pop() (T, bool)`
- `(*PriorityQueue[T]) Peek() (T, bool)`
- `(*PriorityQueue[T]) Len() int`
- `(*PriorityQueue[T]) All() iter.Seq[T]`
- `(*PriorityQueue[T]) Clear()`

Notes:
- Uses `container/heap` internally.
- Behavior (min/max) depends on the `less` comparator.

## OrderedMap[K,V]

- `NewOrderedMap[K,V]() *OrderedMap[K,V]`
- `(*OrderedMap[K,V]) Set(k K, v V)`
- `(*OrderedMap[K,V]) Get(k K) (V, bool)`
- `(*OrderedMap[K,V]) Delete(k K) bool`
- `(*OrderedMap[K,V]) All() iter.Seq2[K,V]`
- `(*OrderedMap[K,V]) Backward() iter.Seq2[K,V]`

## MultiMap[K,V]
- `NewMultiMap[K,V]() *MultiMap[K,V]`
- `(*MultiMap[K,V]) Add(k K, v V)`
- `(*MultiMap[K,V]) Get(k K) []V`
- `(*MultiMap[K,V]) All() iter.Seq2[K,V]`
