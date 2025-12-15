package collections

import "testing"

func TestDequeBasic(t *testing.T) {
	d := NewDeque[int]()
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

func TestDequeEmptyOps(t *testing.T) {
	d := NewDeque[int]()
	if _, ok := d.PeekFront(); ok {
		t.Fatalf("peek front on empty should be false")
	}
	if _, ok := d.PeekBack(); ok {
		t.Fatalf("peek back on empty should be false")
	}
	if _, ok := d.PopFront(); ok {
		t.Fatalf("pop front on empty should be false")
	}
	if _, ok := d.PopBack(); ok {
		t.Fatalf("pop back on empty should be false")
	}
}

func TestDequeClearAndSequence(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushFront(0)
	if d.Len() != 3 {
		t.Fatalf("len after pushes")
	}
	d.Clear()
	if d.Len() != 0 {
		t.Fatalf("len after clear")
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			d.PushFront(i)
		} else {
			d.PushBack(i)
		}
	}
	if d.Len() != 100 {
		t.Fatalf("len after sequence")
	}
	// Drain from front
	for d.Len() > 0 {
		if _, ok := d.PopFront(); !ok {
			t.Fatalf("pop front failed during drain")
		}
	}
	if d.Len() != 0 {
		t.Fatalf("len after drain")
	}
}

func TestDequeCircularBuffer(t *testing.T) {
	d := NewDequeWithCapacity[int](2)
	d.PushBack(1)
	d.PushBack(2)
	if d.Len() != 2 {
		t.Fatalf("len should be 2")
	}
	if v, _ := d.PopFront(); v != 1 {
		t.Fatalf("unexpected pop %d", v)
	}
	d.PushBack(3) // should wrap
	d.PushFront(0)
	if d.Len() != 3 {
		t.Fatalf("len should be 3 after wrap")
	}
	order := []int{}
	for d.Len() > 0 {
		v, _ := d.PopFront()
		order = append(order, v)
	}
	expected := []int{0, 2, 3}
	for i, v := range expected {
		if order[i] != v {
			t.Fatalf("order mismatch %v", order)
		}
	}
}

func TestDequeIterator(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	i := 1
	for v := range d.All() {
		if v != i {
			t.Fatalf("iterator got %d want %d", v, i)
		}
		i++
	}
}

func BenchmarkDequePushPop(b *testing.B) {
	d := NewDeque[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
		d.PopFront()
	}
}

func FuzzDeque(f *testing.F) {
	f.Add([]byte("pushback"))
	f.Fuzz(func(t *testing.T, data []byte) {
		d := NewDeque[byte]()
		for _, b := range data {
			if b%2 == 0 {
				d.PushBack(b)
			} else {
				d.PushFront(b)
			}
		}
		if d.Len() > len(data) {
			t.Errorf("len too big")
		}
		for d.Len() > 0 {
			d.PopFront()
		}
	})
}
