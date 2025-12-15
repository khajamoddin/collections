package itertools

import "iter"

// Map transforms elements of type T to type U.
func Map[T, U any](seq iter.Seq[T], transform func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(transform(v)) {
				return
			}
		}
	}
}

// Filter returns only elements that satisfy the predicate.
func Filter[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if pred(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Reduce accumulates a value over the sequence.
func Reduce[T, Acc any](seq iter.Seq[T], initial Acc, reducer func(Acc, T) Acc) Acc {
	acc := initial
	for v := range seq {
		acc = reducer(acc, v)
	}
	return acc
}

// ToSlice collects the sequence into a slice.
func ToSlice[T any](seq iter.Seq[T]) []T {
	var s []T
	for v := range seq {
		s = append(s, v)
	}
	return s
}
