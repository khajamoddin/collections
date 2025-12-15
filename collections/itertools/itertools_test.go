package itertools

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	seq := func(yield func(int) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}

	transformed := Map(seq, func(v int) int { return v * 2 })
	got := ToSlice(transformed)
	want := []int{2, 4, 6}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	seq := func(yield func(int) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}

	filtered := Filter(seq, func(v int) bool { return v%2 == 0 })
	got := ToSlice(filtered)
	want := []int{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4}
	seq := func(yield func(int) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}

	got := Reduce(seq, 0, func(acc, v int) int { return acc + v })
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestToSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	seq := func(yield func(string) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}

	got := ToSlice(seq)
	if !slices.Equal(got, input) {
		t.Errorf("got %v, want %v", got, input)
	}
}
