package collections

import (
	"reflect"
	"testing"
)

func TestMultiMapBasic(t *testing.T) {
	var mm MultiMap[string, int]
	mm.Add("a", 1)
	mm.Add("a", 2)
	mm.Add("b", 3)

	if got := mm.Len(); got != 3 {
		t.Fatalf("len: got %d", got)
	}
	if !mm.Has("a") || mm.Has("c") {
		t.Fatalf("has check failed")
	}

	vals := mm.Get("a")
	if !reflect.DeepEqual(vals, []int{1, 2}) {
		t.Fatalf("get returned %v", vals)
	}

	if removed := mm.Remove("a", 1); !removed {
		t.Fatalf("remove should succeed")
	}
	if mm.Len() != 2 {
		t.Fatalf("len after remove")
	}

	removedVals := mm.RemoveAll("a")
	if !reflect.DeepEqual(removedVals, []int{2}) {
		t.Fatalf("removeAll returned %v", removedVals)
	}
	if mm.Has("a") {
		t.Fatalf("key should be gone after removeAll")
	}
}

func TestMultiMapKeysValues(t *testing.T) {
	mm := NewMultiMap[int, string]()
	mm.Add(1, "a")
	mm.Add(2, "b")
	mm.Add(2, "c")

	keys := mm.Keys()
	keySet := make(map[int]struct{}, len(keys))
	for _, k := range keys {
		keySet[k] = struct{}{}
	}
	if _, ok := keySet[1]; !ok {
		t.Fatalf("missing key 1")
	}
	if _, ok := keySet[2]; !ok {
		t.Fatalf("missing key 2")
	}

	values := mm.Values()
	if len(values) != 3 {
		t.Fatalf("values length %d", len(values))
	}
	valueSet := make(map[string]int)
	for _, v := range values {
		valueSet[v]++
	}
	if valueSet["c"] != 1 || valueSet["b"] != 1 || valueSet["a"] != 1 {
		t.Fatalf("unexpected values %v", valueSet)
	}

	mm.Clear()
	if mm.Len() != 0 || len(mm.Keys()) != 0 {
		t.Fatalf("clear failed")
	}
}
