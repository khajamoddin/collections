package collections

import "testing"

func TestOrderedMapSetGetDelete(t *testing.T) {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(1, "updated")

	if m.Len() != 2 {
		t.Fatalf("len expected 2, got %d", m.Len())
	}

	if val, ok := m.Get(1); !ok || val != "updated" {
		t.Fatalf("get failed for key 1")
	}
	if val, ok := m.Get(3); ok || val != "" {
		t.Fatalf("expected missing key")
	}

	keys := m.Keys()
	if len(keys) != 2 || keys[0] != 1 || keys[1] != 2 {
		t.Fatalf("keys order incorrect %v", keys)
	}

	values := m.Values()
	if len(values) != 2 || values[0] != "updated" || values[1] != "two" {
		t.Fatalf("values order incorrect %v", values)
	}

	if removed := m.Delete(1); !removed || m.Len() != 1 {
		t.Fatalf("delete failed")
	}
	if m.Has(1) {
		t.Fatalf("key 1 should be gone")
	}
}

func TestOrderedMapRange(t *testing.T) {
	var m OrderedMap[string, int]
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	sum := 0
	m.Range(func(k string, v int) bool {
		sum += v
		return k != "b"
	})
	if sum != 3 { // stopped after b
		t.Fatalf("range early stop failed, sum %d", sum)
	}

	collected := []string{}
	m.RangeReverse(func(k string, v int) bool {
		collected = append(collected, k)
		return true
	})
	if len(collected) != 3 || collected[0] != "c" || collected[2] != "a" {
		t.Fatalf("reverse order incorrect %v", collected)
	}

	m.Clear()
	if m.Len() != 0 || len(m.Keys()) != 0 {
		t.Fatalf("clear failed")
	}
}

func TestOrderedMapDeletionLinks(t *testing.T) {
	m := NewOrderedMap[int, int]()
	for i := 0; i < 5; i++ {
		m.Set(i, i)
	}
	if !m.Delete(0) || !m.Delete(4) || !m.Delete(2) {
		t.Fatalf("delete chain failed")
	}
	want := []int{1, 3}
	got := m.Keys()
	if len(got) != len(want) {
		t.Fatalf("keys length mismatch")
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("keys order mismatch %v", got)
		}
	}
}
