package concurrent_test

import (
	"sync"
	"testing"

	"github.com/khajamoddin/collections/collections/concurrent"
)

func TestSet_Concurrent(t *testing.T) {
	s := concurrent.NewSet[int]()
	var wg sync.WaitGroup
	n := 1000

	// Concurrent writes
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			s.Add(val)
		}(i)
	}
	wg.Wait()

	if s.Len() != n {
		t.Errorf("expected len %d, got %d", n, s.Len())
	}

	// Concurrent reads
	wg = sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			if !s.Has(val) {
				t.Errorf("missing value %d", val)
			}
		}(i)
	}
	wg.Wait()
}

func TestSetWithSnapshot(t *testing.T) {
	s := concurrent.NewSet[int]()
	s.Add(1)
	s.Add(2)

	// Iterator works on snapshot
	count := 0
	for _ = range s.All() {
		count++
		// Modifying set during iteration should not affect iterator
		s.Add(3)
	}

	if count != 2 {
		t.Errorf("expected iterator count 2 (snapshot), got %d", count)
	}
	if s.Len() != 3 {
		t.Errorf("expected final len 3, got %d", s.Len())
	}
}
