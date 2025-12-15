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
for v := range s.All() {
    fmt.Println(v)
}

// Set algebra helpers
other := col.NewSet[int]()
other.Add(30)
union := s.Union(other)           // {10,20,30}
intersection := s.Intersection(s) // {10,20}
fmt.Println(union.Values(), intersection.Values())
```

## Deque[T]

```go
d := col.NewDeque[string]() // uses circular buffer internally
d.PushBack("x")
d.PushFront("y")
front, ok := d.PeekFront()
back, ok2 := d.PeekBack()
_, _ = ok, ok2
v, _ := d.PopFront()
_ = v
d.Clear()
```

## OrderedMap[K,V]

```go
om := col.NewOrderedMap[string, int]()
om.Set("first", 1)
om.Set("second", 2)
v, ok := om.Get("first")
fmt.Println(v, ok)
om.Range(func(k string, v int) bool {
    fmt.Println(k, v) // visits in insertion order
    return true
})

// Or use Go 1.23 iterator
for k, v := range om.All() {
    fmt.Println(k, v)
}
keys := om.Keys()   // ["first", "second"]
vals := om.Values() // [1, 2]
_ = om.Delete("first")
```

## MultiMap[K,V]

```go
mm := col.NewMultiMap[string, int]()
mm.Add("id", 1)
mm.Add("id", 2)
fmt.Println(mm.Get("id")) // [1 2]
mm.Remove("id", 1)
mm.RemoveAll("id") // returns removed values
fmt.Println(mm.Len())
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

- `examples/set/main.go`
- `examples/deque/main.go`
- `examples/iter/iter_example.go` (shows new Go 1.23 iterators)
