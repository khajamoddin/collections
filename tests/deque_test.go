package collections_test

import (
    col "github.com/khajamoddin/collections/collections"
    "testing"
)

func TestDequeBasic(t *testing.T) {
    d := col.NewDeque[int]()
    if d.Len() != 0 {
        t.Fatalf("len")
    }
    d.PushBack(1)
    d.PushFront(0)
    if d.Len() != 2 {
        t.Fatalf("len")
    }
    v, ok := d.PeekFront()
    if !ok || v != 0 {
        t.Fatalf("peek front")
    }
    v, ok = d.PopFront()
    if !ok || v != 0 {
        t.Fatalf("pop front")
    }
    v, ok = d.PopBack()
    if !ok || v != 1 {
        t.Fatalf("pop back")
    }
    if d.Len() != 0 {
        t.Fatalf("len")
    }
}

