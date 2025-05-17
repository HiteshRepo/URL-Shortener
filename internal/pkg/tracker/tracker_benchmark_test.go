package tracker

import (
	"math/rand"
	"testing"
)

var domains = []string{
	"google.com",
	"github.com",
	"stackoverflow.com",
	"google.com",
	"github.com",
	"example.com",
	"amazon.com",
	"google.com",
	"netflix.com",
	"netflix.com",
	"netflix.com",
	"netflix.com",
}

func Benchmark_MaxHeap_GetTopN(b *testing.B) {
	heap := setupHeap(1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.GetTopN(10)
	}
}

func Benchmark_MaxHeap_Visit(b *testing.B) {
	heap := NewMaxHeap()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randomIndex := rand.Intn(len(domains))
		selectedDomain := domains[randomIndex]

		heap.Visit(selectedDomain)
	}
}

func Benchmark_BruteForce_GetTopN(b *testing.B) {
	tracker := setupTracker(1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tracker.GetTopN(10)
	}
}

func Benchmark_BruteForce_Visit(b *testing.B) {
	tracker := NewBruteForceTracker()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randomIndex := rand.Intn(len(domains))
		selectedDomain := domains[randomIndex]

		tracker.Visit(selectedDomain)
	}
}

func setupHeap(size int) *MaxHeap {
	heap := NewMaxHeap()

	for i := 0; i < size; i++ {
		randomIndex := rand.Intn(len(domains))
		selectedDomain := domains[randomIndex]

		heap.Visit(selectedDomain)
	}

	return heap
}

func setupTracker(size int) *BruteForceTracker {
	tracker := NewBruteForceTracker()

	for i := 0; i < size; i++ {
		randomIndex := rand.Intn(len(domains))
		selectedDomain := domains[randomIndex]

		tracker.Visit(selectedDomain)
	}

	return tracker
}
