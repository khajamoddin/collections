// Package itertools provides functional helpers for working with Go 1.23+ iterators.
//
// It contains standard operations like Map, Filter, and Reduce that operate directly
// on iter.Seq[T], allowing for expressive and concise data transformation pipelines.
//
// Example:
//
//	ints := slices.Values([]int{1, 2, 3})
//	doubled := itertools.Map(ints, func(x int) int { return x * 2 })
//	list := itertools.ToSlice(doubled) // [2, 4, 6]
package itertools
