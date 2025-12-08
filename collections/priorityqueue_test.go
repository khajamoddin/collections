package collections

import "testing"

func TestPriorityQueueBasic(t *testing.T) {
	q := NewPriorityQueue[int](func(a, b int) bool { return a < b })
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

func TestPriorityQueueMaxHeap(t *testing.T) {
    q := NewPriorityQueue[int](func(a, b int) bool { return a > b })
    for _, v := range []int{5, 1, 3, 9, 7} {
        q.Push(v)
    }
    prev := 1<<31 - 1
    for q.Len() > 0 {
        v, ok := q.Pop()
        if !ok {
            t.Fatalf("pop failed")
        }
        if v > prev { // should be non-increasing sequence
            t.Fatalf("sequence not max-heap order")
        }
        prev = v
    }
}

func TestPriorityQueueStress(t *testing.T) {
    q := NewPriorityQueue[int](func(a, b int) bool { return a < b })
    for i := 1000; i >= 1; i-- {
        q.Push(i)
    }
    expected := 1
    for q.Len() > 0 {
        v, ok := q.Pop()
        if !ok || v != expected {
            t.Fatalf("expected %d, got %d", expected, v)
        }
        expected++
    }
}
