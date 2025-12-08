---
layout: default
title: Usage Guide
nav_order: 3
---

# Usage Guide

## Import

```go
import col "github.com/khajamoddin/collections/collections"
```

## Set[T]

```go
s := col.NewSet[int]()
s.Add(10)
s.Add(20)
_ = s.Has(10)
s.Remove(20)
fmt.Println(s.Len())
for _, v := range s.Values() {
    fmt.Println(v)
}
```

## Deque[T]

```go
d := col.NewDeque[string]()
d.PushBack("x")
d.PushFront("y")
front, ok := d.PeekFront()
back, ok2 := d.PeekBack()
_, _ = ok, ok2
v, _ := d.PopFront()
_ = v
d.Clear()
```

## PriorityQueue[T]

```go
pq := col.NewPriorityQueue[int](func(a, b int) bool { return a < b })
pq.Push(3)
pq.Push(1)
pq.Push(2)
peek, _ := pq.Peek()
fmt.Println(peek) // 1
for pq.Len() > 0 {
    v, _ := pq.Pop()
    fmt.Println(v)
}
```

## Examples Folder

- `examples/set_example.go` is runnable.
- `examples/deque_example.go` contains `dequeExample()` to demonstrate operations.
