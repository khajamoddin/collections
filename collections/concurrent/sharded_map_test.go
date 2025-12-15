package concurrent_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/khajamoddin/collections/collections/concurrent"
)

func TestShardedMap_Concurrent(t *testing.T) {
	m := concurrent.NewShardedMap[string, int](16, concurrent.StringHasher)
	var wg sync.WaitGroup
	n := 1000

	// Concurrent writes
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			m.Set(strconv.Itoa(val), val)
		}(i)
	}
	wg.Wait()

	if m.Len() != n {
		t.Errorf("expected len %d, got %d", n, m.Len())
	}

	// Concurrent reads
	wg = sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			v, ok := m.Get(strconv.Itoa(val))
			if !ok || v != val {
				t.Errorf("missing or incorrect value for %d", val)
			}
		}(i)
	}
	wg.Wait()
}

func TestShardedMap_Snapshot(t *testing.T) {
	m := concurrent.NewShardedMap[string, int](4, concurrent.StringHasher)
	m.Set("a", 1)
	m.Set("b", 2)

	// Iterator works on snapshot
	count := 0
	for _, v := range m.All() {
		count++
		// Modification during iteration should validly proceed but not affect current iterator
		m.Set("c", 3+v)
	}

	if count != 2 {
		t.Errorf("expected iterator count 2, got %d", count)
	}
	if m.Len() != 3 {
		t.Errorf("expected final len 3, got %d", m.Len())
	}
}
