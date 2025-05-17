package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopN(t *testing.T) {
	tracker := NewBruteForceTracker()

	domains := []string{
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
	expected := []Domain{
		{Name: "netflix.com", Count: 4},
		{Name: "google.com", Count: 3},
		{Name: "github.com", Count: 2},
	}

	for _, domain := range domains {
		tracker.Visit(domain)
	}

	top3 := tracker.GetTopN(3)
	assert.ElementsMatch(t, expected, top3)
}
