package collections_test

import (
    col "github.com/khajamoddin/collections/collections"
    "testing"
)

func TestSetBasic(t *testing.T) {
    s := col.NewSet[int]()
    if s.Len() != 0 {
        t.Fatalf("len")
    }
    s.Add(1)
    s.Add(2)
    s.Add(2)
    if !s.Has(2) {
        t.Fatalf("has")
    }
    if s.Len() != 2 {
        t.Fatalf("len")
    }
    s.Remove(2)
    if s.Has(2) {
        t.Fatalf("remove")
    }
    s.Clear()
    if s.Len() != 0 {
        t.Fatalf("clear")
    }
}

func BenchmarkSetAdd(b *testing.B) {
    s := col.NewSet[int]()
    for i := 0; i < b.N; i++ {
        s.Add(i)
    }
}

