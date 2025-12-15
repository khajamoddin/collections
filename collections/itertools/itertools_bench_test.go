package itertools

import (
	"testing"

	"iter"
)

// rangeSeq returns a sequence [0, n).
func rangeSeq(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func BenchmarkMap_IntToInt_Iter(b *testing.B) {
	const n = 100_000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seq := Map(rangeSeq(n), func(v int) int {
			return v * 2
		})

		// Force evaluation
		sum := 0
		for v := range seq {
			sum += v
		}
		if sum == 0 {
			b.Fatalf("unexpected sum = 0")
		}
	}
}

func BenchmarkMap_IntToInt_Slice(b *testing.B) {
	const n = 100_000

	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range src {
			sum += v * 2
		}
		if sum == 0 {
			b.Fatalf("unexpected sum = 0")
		}
	}
}

func BenchmarkFilter_Even_Iter(b *testing.B) {
	const n = 100_000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seq := Filter(rangeSeq(n), func(v int) bool {
			return v%2 == 0
		})

		count := 0
		for range seq {
			count++
		}
		if count == 0 {
			b.Fatalf("unexpected count = 0")
		}
	}
}

func BenchmarkFilter_Even_Slice(b *testing.B) {
	const n = 100_000

	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		for _, v := range src {
			if v%2 == 0 {
				count++
			}
		}
		if count == 0 {
			b.Fatalf("unexpected count = 0")
		}
	}
}

func BenchmarkReduce_Sum_Iter(b *testing.B) {
	const n = 100_000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := Reduce(rangeSeq(n), 0, func(acc, v int) int {
			return acc + v
		})
		if sum == 0 {
			b.Fatalf("unexpected sum = 0")
		}
	}
}

func BenchmarkReduce_Sum_Slice(b *testing.B) {
	const n = 100_000

	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range src {
			sum += v
		}
		if sum == 0 {
			b.Fatalf("unexpected sum = 0")
		}
	}
}

func BenchmarkMapFilter_Pipeline_Iter(b *testing.B) {
	const n = 100_000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// (i * 2) where (i * 2) % 3 == 0
		seq := Map(
			Filter(rangeSeq(n), func(v int) bool {
				return v%2 == 0
			}),
			func(v int) int {
				return v * 2
			},
		)

		sum := 0
		for v := range seq {
			if v%3 == 0 {
				sum += v
			}
		}
		if sum == 0 {
			b.Fatalf("unexpected sum = 0")
		}
	}
}

func BenchmarkMapFilter_Pipeline_Slice(b *testing.B) {
	const n = 100_000

	src := make([]int, n)
	for i := 0; i < n; i++ {
		src[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range src {
			if v%2 == 0 {
				v2 := v * 2
				if v2%3 == 0 {
					sum += v2
				}
			}
		}
		if sum == 0 {
			b.Fatalf("unexpected sum = 0")
		}
	}
}
