# Tracker Package

The `tracker` package provides utilities for tracking domain visit counts and efficiently retrieving the top N most visited domains. It offers two implementations with different performance characteristics: a brute-force tracker and a max-heap-based tracker.

## Features

- Track the number of visits to each domain.
- Retrieve the top N most visited domains.
- Two tracker implementations:
  - **BruteForceTracker**: Simple, map-based approach.
  - **MaxHeap**: Efficient heap-based approach for fast top-N queries.
- Comprehensive unit tests and benchmarks.

## Directory Structure

```
internal/pkg/tracker/
├── brute_force.go             # BruteForceTracker implementation
├── brute_force_test.go        # Tests for BruteForceTracker
├── domain.go                  # Domain struct and interface
├── max_heap.go                # MaxHeap tracker implementation
├── max_heap_test.go           # Tests for MaxHeap
└── tracker_benchmark_test.go  # Benchmarks for both trackers
```

## Usage

### Domain Struct

All trackers operate on the following `Domain` struct:

```go
type Domain struct {
    Name  string
    Count int
}
```

### BruteForceTracker

A simple tracker using a map to count visits.

```go
import "internal/pkg/tracker"

tracker := tracker.NewBruteForceTracker()
tracker.Visit("google.com")
tracker.Visit("github.com")
tracker.Visit("google.com")

top := tracker.GetTopN(2) // Returns the top 2 most visited domains
fmt.Println(top) // [{google.com 2} {github.com 1}]
fmt.Println(tracker.Stats()) // Total visits: 3, Unique domains: 2
```

### MaxHeap

A tracker using a max-heap for efficient retrieval of top N domains.

```go
import "internal/pkg/tracker"

heap := tracker.NewMaxHeap()
heap.Visit("google.com")
heap.Visit("github.com")
heap.Visit("google.com")

top := heap.GetTopN(2) // Returns the top 2 most visited domains
fmt.Println(top) // [{google.com 2} {github.com 1}]
```

## Interface

The `DomainTrackerer` interface (see `domain.go`) can be implemented for custom tracking strategies:

```go
type DomainTrackerer interface {
    GetTopN() []Domain
    Visit(domainName string)
}
```

## Testing

Unit tests are provided for both tracker implementations:

- `brute_force_test.go`
- `max_heap_test.go`

Run all tests:

```sh
go test ./internal/pkg/tracker/...
```

## Benchmarking

Benchmarks compare the performance of both trackers for `Visit` and `GetTopN` operations. See `tracker_benchmark_test.go`.

Run benchmarks:

```sh
go test -bench=. ./internal/pkg/tracker
```

## Extensibility

You can add new tracker implementations by implementing the `DomainTrackerer` interface.

## License

This package is part of the URL-Shortener project.
