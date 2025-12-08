# Testing & Benchmarks

## Run Tests

```bash
go test ./... -v -cover
```

## Benchmarks

Example benchmark for `Set[T]`:

```go
func BenchmarkSetAdd(b *testing.B) {
    s := NewSet[int]()
    for i := 0; i < b.N; i++ {
        s.Add(i)
    }
}
```

Run:

```bash
go test -bench=. ./...
```

## Coverage

To improve coverage, place tests in the same package directory (e.g., `collections/collections/*_test.go`) so the coverage tool instruments the package under test.
