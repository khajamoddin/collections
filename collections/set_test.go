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

func TestSetOperations(t *testing.T) {
	a := NewSet[int]()
	b := NewSet[int]()
	for _, v := range []int{1, 2, 3} {
		a.Add(v)
	}
	for _, v := range []int{3, 4} {
		b.Add(v)
	}

	if !a.IsSuperset(a.Intersection(a)) {
		t.Fatalf("intersection subset check failed")
	}

	u := a.Union(b)
	for _, v := range []int{1, 2, 3, 4} {
		if !u.Has(v) {
			t.Fatalf("union missing %d", v)
		}
	}

	i := a.Intersection(b)
	if i.Len() != 1 || !i.Has(3) {
		t.Fatalf("intersection incorrect")
	}

	d := a.Difference(b)
	if d.Len() != 2 || d.Has(3) {
		t.Fatalf("difference incorrect")
	}

	sd := a.SymmetricDifference(b)
	for _, v := range []int{1, 2, 4} {
		if !sd.Has(v) {
			t.Fatalf("symmetric difference missing %d", v)
		}
	}

	other := NewSet[int]()
	other.Add(5)
	if !b.IsDisjoint(other) {
		t.Fatalf("disjointness failed")
	}

	if !a.IsSubset(u) || !u.IsSuperset(a) {
		t.Fatalf("subset/superset failed")
	}

	clone := a.Clone()
	a.Add(99)
	if clone.Has(99) || clone.Len() != 3 {
		t.Fatalf("clone should be independent")
	}
}

func TestSetNilSafety(t *testing.T) {
	var s *Set[int]
	if s.Len() != 0 || s.Has(1) {
		t.Fatalf("nil set should behave empty")
	}
	s.Add(1) // should not panic
	if s != nil {
		t.Fatalf("nil set should remain nil after Add")
	}
	clone := s.Clone()
	if clone.Len() != 0 {
		t.Fatalf("clone of nil should be empty")
	}
}
