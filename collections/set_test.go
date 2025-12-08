package collections

import "testing"

func TestSetBasic(t *testing.T) {
    s := NewSet[int]()
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
    s := NewSet[int]()
    for i := 0; i < b.N; i++ {
        s.Add(i)
    }
}

func TestSetValuesOrderNonDeterministic(t *testing.T) {
    s := NewSet[int]()
    for i := 0; i < 10; i++ {
        s.Add(i)
    }
    vals := s.Values()
    if len(vals) != s.Len() {
        t.Fatalf("values length mismatch")
    }
    seen := make(map[int]bool)
    for _, v := range vals {
        seen[v] = true
    }
    for i := 0; i < 10; i++ {
        if !seen[i] {
            t.Fatalf("missing value %d", i)
        }
    }
}
