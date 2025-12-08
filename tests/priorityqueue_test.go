package collections_test

import (
    col "github.com/khajamoddin/collections/collections"
    "testing"
)

func TestPriorityQueueBasic(t *testing.T) {
    q := col.NewPriorityQueue[int](func(a, b int) bool { return a < b })
    q.Push(3)
    q.Push(1)
    q.Push(2)
    v, ok := q.Peek()
    if !ok || v != 1 {
        t.Fatalf("peek")
    }
    v, ok = q.Pop()
    if !ok || v != 1 {
        t.Fatalf("pop 1")
    }
    v, ok = q.Pop()
    if !ok || v != 2 {
        t.Fatalf("pop 2")
    }
    v, ok = q.Pop()
    if !ok || v != 3 {
        t.Fatalf("pop 3")
    }
    _, ok = q.Pop()
    if ok {
        t.Fatalf("pop empty")
    }
}

